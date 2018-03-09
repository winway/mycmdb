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

</head>

<body>
  <form id="app" class="form-horizontal" enctype="multipart/form-data" method="post" action="">

      <div class="form-group">
        <label for="Content" class="col-sm-3 control-label">说明：</label>
        <div class="col-sm-5">
          <textarea v-model="Content" class="form-control" style="height:300px" name="Content" placeholder="请填写相关信息" readonly>{{ .step.Content }}</textarea>
        </div>
      </div>

      <div class="form-group">
        <label for="Content" class="col-sm-3 control-label">附件：</label>
        <div class="col-sm-5">
          <a href="/static/upload/{{ .step.File }}">{{ .step.File }}</a>
        </div>
      </div>

  </form>
</body>
</html>
