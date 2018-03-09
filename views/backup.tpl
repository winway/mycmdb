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
  <link href="/static/css/custom.css" rel="stylesheet">

  <style>
  </style>
</head>

<body>
  <div class="wrapper">
    <!-- Sidebar Holder -->
    <nav id="sidebar">
      <div class="sidebar-header">
        <h3>运维自动化平台</h3>
        <strong><img src="/static/img/favicon.ico" /></strong>
      </div>

      <ul class="list-unstyled components">
        <li class="active">
          <a href="#homeSubmenu" data-toggle="collapse" aria-expanded="false">
            <i class="glyphicon glyphicon-folder-open"></i>
            资源管理
          </a>
          <ul class="collapse list-unstyled" id="homeSubmenu">
            <li><a href="#">机房管理</a></li>
            <li><a href="#">IP管理</a></li>
            <li><a href="#">服务器管理</a></li>
          </ul>
        </li>

        <li>
          <a href="#">
            <i class="glyphicon glyphicon-wrench"></i>
            OS安装
          </a>
        </li>

        <li>
          <a href="#">
            <i class="glyphicon glyphicon-share-alt"></i>
            代码发布
          </a>
        </li>

        <li>
          <a href="#pageSubmenu" data-toggle="collapse" aria-expanded="false">
            <i class="glyphicon glyphicon-user"></i>
            用户管理
          </a>
          <ul class="collapse list-unstyled" id="pageSubmenu">
            <li><a href="#">用户管理</a></li>
            <li><a href="#">组管理</a></li>
            <li><a href="#">角色管理</a></li>
            <li><a href="#">权限管理</a></li>
          </ul>
        </li>
      </ul>

      <ul class="list-unstyled CTAs">
        <li><a href="#" class="download">监控系统</a></li>
        <li><a href="#" class="article">日志系统</a></li>
      </ul>
    </nav>

    <!-- Page Content Holder -->
    <div id="content" class="container" style="width:100%">

      <nav class="navbar navbar-default">
        <div class="container-fluid">
          <div class="navbar-header">
            <button type="button" id="sidebarCollapse" class="btn btn-info navbar-btn">
              <i class="glyphicon glyphicon-align-left"></i>
              <span></span>
            </button>
          </div>

          <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul class="nav navbar-nav navbar-right">
              <button class="btn btn-success" type="button">
                通知 <span class="badge">4</span>
              </button>
              <!-- Split button -->
              <div class="btn-group navbar-btn">
                <button type="button" class="btn btn-default">注销</button>
                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                  <span class="caret"></span>
                  <span class="sr-only">Toggle Dropdown</span>
                </button>
                <ul class="dropdown-menu">
                  <li><a href="#">设置</a></li>
                  <li role="separator" class="divider"></li>
                  <li><a href="#">修改密码</a></li>
                </ul>
              </div>
            </ul>
          </div>
        </div>
      </nav>

      <h2>Coming soon!</h2>

    </div>
  </div>

  <!-- Placed at the end of the document so the pages load faster -->
  <script src="/static/js/jquery-3.2.1.min.js"></script>
  <script src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>

  <script>
    $(document).ready(function () {
      $('#sidebarCollapse').on('click', function () {
        $('#sidebar').toggleClass('active');
      });
    });
  </script>
</body>

</html>