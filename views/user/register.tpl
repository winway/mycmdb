<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="utf-8">
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="运维自动化平台">
  <meta name="keywords" content="devops">
  <meta name="author" content="wangwei03@sunlands.com">
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
  {{ else if .flash.notice }}
  <div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
  {{ end }}

  <div id="app" class="container">
    <form class="form-horizontal" method="post" action="">
      <div class="form-group">
        <label for="email" class="col-sm-2 control-label">邮箱：</label>
        <div class="col-sm-10">
          <input v-model="email" type="email" class="form-control" name="email" placeholder="请输入邮箱" required>
        </div>
      </div>

      <div class="form-group">
        <label for="name" class="col-sm-2 control-label">姓名：</label>
        <div class="col-sm-10">
          <input v-model="name" type="text" class="form-control" name="name" placeholder="请输入姓名" required>
        </div>
      </div>

      <div class="form-group">
        <label for="password" class="col-sm-2 control-label">密码：</label>
        <div class="col-sm-10">
          <input v-model="password" type="password" class="form-control" name="password" placeholder="请输入密码" required>
        </div>
      </div>

      <div class="form-group">
        <label for="password2" class="col-sm-2 control-label">确认密码：</label>
        <div class="col-sm-10">
          <input v-model="password2" type="password" class="form-control" name="password2" placeholder="请输入密码" required>
        </div>
      </div>

      <div>
        <div class="col-sm-2"></div>
        <button type="button" v-on:click="save" class="btn btn-primary col-sm-10">保存</button>
      </div>
    </form>
  </div> <!-- /container -->

  <script type="text/javascript">
    var vue = new Vue({
      el: '#app',
      data: {
        email: "",
        name: "",
        password: "",
        password2: ""
      },
      methods: {
        save: function() {
          var self = this;
          var emailRE = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

          if (!emailRE.test(self.email)) {
            alert("邮箱格式不合法");
            return false;
          }

          if (self.name == "" || self.password == "") {
            alert("信息不全");
            return false;
          }

          if (self.password != self.password2) {
            alert("密码不一致");
            return false;
          }

          $.ajax({
            type:"POST",
            url:"{{ urlfor "UserController.Register" }}",
            data:{
              'email': self.email,
              'name': self.name,
              'password': self.password
            },
            success: function(data, textStatus, jqXHR) {
              if (data.code != 0) {
                alert(data.msg);
              } else {
                window.location.href = "{{ urlfor "MainController.Get" }}";
              }
            }
          });

        }
      }
    });
  </script>
</body>
</html>
