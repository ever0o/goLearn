<html>
<head>
    <title></title>
</head>
<body>
<form action="/login" method="post">
    <table>
        <tr>
            <td>
                用户名:
            </td>
            <td>
                <input type="text" name="username"/>
            </td>
        </tr>
        <tr>
            <td>
                真实姓名:
            </td>
            <td>
                <input type="text" name="realName"/>
            </td>
        </tr>
        <tr>
            <td>
                英文名称:
            </td>
            <td>
                <input type="text" name="engName"/>
            </td>
        </tr>
        <tr>
            <td>
                邮件地址:
            </td>
            <td>
                <input type="text" name="email"/>
            </td>
        </tr>
        <tr>
            <td>
                手机号码:
            </td>
            <td>
                <input type="text" name="phone"/>
            </td>
        </tr>
        <tr>
            <td>
                密码:
            </td>
            <td>
                <input type="password" name="password"/>
            </td>
        </tr>
        <tr>
            <td>
                年龄:
            </td>
            <td>
                <input type="text" name="age"/>
            </td>
        </tr>
        <tr>
            <td>
                水果选择:
            </td>
            <td>
                <select name="fruit">
                    <option value="apple">apple</option>
                    <option value="banana">banana</option>
                    <option value="pear">pear</option>
                </select>
            </td>
        </tr>
        <tr>
            <td>
                性别选择:
            </td>
            <td>
                <input type="radio" name="gender" value="1"/>男
            </td>
            <td>
                <input type="radio" name="gender" value="2"/>女
            </td>
        </tr>
        <tr>
            <td>
                运动选择:
            </td>
            <td>
                <input type="checkbox" name="interest" value="football"/>足球
                <input type="checkbox" name="interest" value="basketball"/>篮球
                <input type="checkbox" name="interest" value="tennis"/>网球
            </td>
        </tr>
        <tr>
            <td>
                <input type="hidden" name="token" value="{{.}}"/>
            </td>
        </tr>
        <tr>
            <td>
                <input type="submit" value="登陆"/>
            </td>
        </tr>
    </table>
</form>
</body>
</html>