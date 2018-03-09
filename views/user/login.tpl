<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="utf-8">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="运维自动化平台">
  <meta name="keywords" content="devops">
  <meta name="author" content="winway1988@163.com">
  <meta name="robots" content="index,follow">

  <title>运维自动化平台</title>

  <link href="/static/img/favicon.ico" rel="shortcut icon">

  <!-- Site CSS -->
  <link href="/static/bootstrap-3.3.7-dist/css/bootstrap.min.css" rel="stylesheet">
  <link href="/static/bootstrap-3.3.7-dist/css/bootstrap-theme.min.css" rel="stylesheet">
  <link href="/static/DataTables-1.10.15/media/css/jquery.dataTables.min.css" rel="stylesheet">
  <link href="/static/DataTables-1.10.15/media/css/dataTables.bootstrap.min.css" rel="stylesheet">
  <link href="/static/bootstrap-select-1.12.4/dist/css/bootstrap-select.min.css" rel="stylesheet">

  <!-- Placed at the end of the document so the pages load faster -->
  <script type="text/javascript" src="/static/jquery-1.12.4/dist/jquery.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/DataTables-1.10.15/media/js/jquery.dataTables.min.js"></script>
  <script type="text/javascript" src="/static/DataTables-1.10.15/media/js/dataTables.bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-select-1.12.4/dist/js/bootstrap-select.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-select-1.12.4/dist/js/i18n/defaults-zh_CN.min.js"></script>
  <script type="text/javascript" src="/static/js/vue.min.js"></script>

  <style>
  body {
    padding-top: 40px;
    padding-bottom: 40px;
    background-color: #eee;
  }

  .form-signin {
    max-width: 330px;
    padding: 15px;
    margin: 0 auto;
  }
  .form-signin .form-signin-heading,
  .form-signin .checkbox {
    margin-bottom: 10px;
  }
  .form-signin .checkbox {
    font-weight: 400;
  }
  .form-signin .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
  }
  .form-signin .form-control:focus {
    z-index: 2;
  }
  .form-signin input[type="email"] {
    margin-bottom: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
  }
  .form-signin input[type="password"] {
    margin-bottom: 10px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
  }


</style>
</head>

<body>
  {{ if .flash.error }}
  <div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
  {{ end }}

  <div class="container">
    <form class="form-signin" method="post" action="">
      <h2 class="form-signin-heading text-center"><img src="/static/img/favicon.ico"/>&nbsp;&nbsp;&nbsp;运维自动化平台</h2>
      <input type="hidden" name="rf" placeholder="" value="{{ .referer }}">
      <label for="inputEmail" class="sr-only">邮箱地址</label>
      <input type="email" id="inputEmail" name="email" class="form-control" placeholder="Email address" required autofocus>
      <label for="inputPassword" class="sr-only">密码</label>
      <input type="password" id="inputPassword" name="password" class="form-control" placeholder="Password" required>
      <div class="checkbox">
        <label>
          <input type="checkbox" value="remember-me"> Remember me
        </label>
      </div>

      <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
      <a href="{{ urlfor "UserController.RegisterPage" }}">Register</a>
    </form>
  </div> <!-- /container -->
</body>
</html>
