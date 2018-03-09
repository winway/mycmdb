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
  <link href="/static/css/custom.css" rel="stylesheet">

  <!-- Placed at the end of the document so the pages load faster -->
  <script type="text/javascript" src="/static/jquery-1.12.4/dist/jquery.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/DataTables-1.10.15/media/js/jquery.dataTables.min.js"></script>
  <script type="text/javascript" src="/static/DataTables-1.10.15/media/js/dataTables.bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-select-1.12.4/dist/js/bootstrap-select.min.js"></script>
  <script type="text/javascript" src="/static/bootstrap-select-1.12.4/dist/js/i18n/defaults-zh_CN.min.js"></script>
  <script type="text/javascript" src="/static/js/vue.min.js"></script>
  <script type="text/javascript" src="/static/js/custom.js"></script>

  {{template "head" .}}
</head>

<body>
  <div class="wrapper">
    <!-- Sidebar Holder -->
    {{template "sidebar.tpl" .}}

    <!-- Page Content Holder -->
    <div id="content" class="container" style="width:100%">
      {{template "navbar.tpl" .}}

      {{template "body" .}}

    </div>
  </div>

  <script>
    $(document).ready(function () {
      {{ if .sbStatus }}
      $('#sidebar').toggleClass('active');
      {{ end  }}

      {{ if .umStatus }}
      $('#userManage').trigger("click");
      {{ end }}

      {{ if .smStatus }}
      $('#SourceManage').trigger("click");
      {{ end }}

      $('#sidebarCollapse').on('click', function () {
        $('#sidebar').toggleClass('active');

        $.ajax({
          type: "POST",
          url: "/user/menu/?type=sb"
        }).done(function (data) {
        });
      });

      $('#userManage').on('click', function () {
        $.ajax({
          type: "POST",
          url: "/user/menu/?type=um"
        }).done(function (data) {
        });
      });

      $('#SourceManage').on('click', function () {
        $.ajax({
          type: "POST",
          url: "/user/menu/?type=sm"
        }).done(function (data) {
        });
      });

    });
  </script>
</body>

</html>