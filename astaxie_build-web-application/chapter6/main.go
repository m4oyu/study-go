package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"study-go/astaxie_build-web-application/chapter6/session"
	_ "study-go/astaxie_build-web-application/chapter6/session/memory"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
	//注意：もしParseFormメソッドがコールされなければ、以下でフォームのデータを取得することができません。
	fmt.Println(r.Form) //これらのデータはサーバのプリント情報に出力されます
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //ここでwに書き込まれたものがクライアントに出力されます。
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/count", 302)
	}
}

//func count(w http.ResponseWriter, r *http.Request) {
//	sess := globalSessions.SessionStart(w, r)
//	createtime := sess.Get("createtime")
//	if createtime == nil {
//		sess.Set("createtime", time.Now().Unix())
//	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
//		globalSessions.SessionDestroy(w, r)
//		sess = globalSessions.SessionStart(w, r)
//	}
//	ct := sess.Get("countnum")
//	if ct == nil {
//		sess.Set("countnum", 1)
//	} else {
//		sess.Set("countnum", ct.(int) + 1)
//	}
//	t, _ := template.ParseFiles("count.gtpl")
//	w.Header().Set("Content-Type", "text/html")
//	t.Execute(w, sess.Get("countnum"))
//}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", ct.(int)+1)
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

var globalSessions *session.Manager

//この後、init関数で初期化を行います。
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/", sayHelloName) //アクセスのルーティングを設定します
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)

	//アクセスのルーティングを設定します
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
