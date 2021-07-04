<html>
<head>
    <title></title>
</head>
<body>
<form action="/login?username=attaxie" method="post">
    ユーザ名:<input type="text" name="username">
    パスワード:<input type="password" name="password">
    <input type="submit" value="ログイン">
</form>

<form action="/pullDown" method="post">
    <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banana">banana</option>
    </select>
    <input type="submit" value="submit">
</form>
</body>
</html>