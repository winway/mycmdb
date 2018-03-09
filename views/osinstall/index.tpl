{{template "base.tpl" .}}

{{define "head"}}
{{end}}

{{define "body"}}

<div class="page-header">
  <h2>OS自助安装&nbsp;&nbsp;&nbsp;&nbsp;<small>累计安装服务器{{ .doneCnt }}台 正在安装{{ .doingCnt }}台</small></h2>
</div>

{{ if .flash.error }}
<div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
{{ else if .flash.notice }}
<div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
{{ end }}

<div>
  <form action="" class="form-inline">
    <div class="form-group">
      <label  class="control-label" for="InputIp">IP:</label>
      <input id="InputIp" class="form-control" placeholder="请输入IP">

      <label  class="control-label" for="selOsVersion">OS:</label>
      <select id="selOsVersion" multiple class="selectpicker" title="请选择OS" data-style="btn-default">
        <option value=0>CentOS-6.9-x86_64</option>
        <option value=1>centos73-x86_64</option>
      </select>

      <label  class="control-label" for="selStatus">状态:</label>
      <select id="selStatus" multiple class="selectpicker" title="请选择状态" data-style="btn-default">
        <option value=0>未开始</option>
        <option value=1>正在安装</option>
        <option value=2>完成</option>
      </select>

      <a onclick="Search()" class="btn btn-primary"><i class="glyphicon glyphicon-search"></i>查询</a>

      <a class="btn btn-primary" href="{{ urlfor "OsInstallController.ApplyPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-retweet"></i>安装申请
      </a>

      <a onclick="Delete('{{ urlfor "OsInstallController.Cancel" }}')" class="btn btn-primary"><i class="glyphicon glyphicon-remove-circle"></i>取消安装</a>
    </div>
  </form>
</div>


<!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog">
   <div class="modal-content">
   </div>
 </div><!-- /.modal -->
</div>


<div>
  <table id="myTable"  class="display" cellpadding="0" cellspacing="0" width="100%">
    <thead>
      <tr>
        <th><input type="checkbox" name="checkall" id="checkall"></th>
        <th>ID</th>
        <th>SN</th>
        <th>远程卡IP</th>
        <th>网卡IP</th>
        <th>OS</th>
        <th>状态</th>
        <th>创建时间</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
    </tbody>
  </table>
</div>


<script type="text/javascript">
  var table;
  $(document).ready(function() {
    table = $('#myTable').dataTable({
    "sPaginationType": "full_numbers", // 分页风格，full_number会把所有页码显示出来（大概是，自己尝试）
    "iDisplayLength": 10,
    "bAutoWidth": false,
    "bLengthChange": false,
    "bFilter": false,
    "bSort": true,
    "oLanguage": {    //下面是一些汉语翻译
      "sSearch": "搜索",
      "sLengthMenu": "每页显示 _MENU_ 条记录",
      "sZeroRecords": "没有检索到数据",
      "sInfo": "显示 _START_ 至 _END_ 条 &nbsp;&nbsp;共 _TOTAL_ 条",
      "sInfoFiltered": "(筛选自 _MAX_ 条数据)",
      "sInfoEmtpy": "没有数据",
      //"sProcessing": "正在加载数据...",
      "sProcessing": "<img src='/static/img/table_loading.gif' />",
      "oPaginate":
      {
        "sFirst": "首页",
        "sPrevious": "前一页",
        "sNext": "后一页",
        "sLast": "末页"
      }
    },
    "bProcessing": true,
    "bServerSide": true,
    "sAjaxSource": "{{ urlfor "OsInstallController.List" }}",
    "aoColumns": [
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<input id=' + full['Id'] + ' value=' + full['Id'] + ' name="checksingle" type="checkbox">';
      }
    },
    {"mData": 'Id', "sClass": "alignRight"},
    {"mData": 'Server.Sn', "sClass": "alignRight"},
    {"mData": 'Server.RemoteCardIp', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return 'Eth1: ' + full['Server']['Eth1Ip'] + '<br />' + 'Eth2: ' + full['Server']['Eth2Ip'] + '<br />' + 'Eth3: ' + full['Server']['Eth3Ip'] + '<br />' + 'Eth4: ' + full['Server']['Eth4Ip'];
      }
    },
    {
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        if (full["Server"]["OsVersion"] == 0) {
          return "CentOS-6.9-x86_64";
        } else if (full["Server"]["OsVersion"] == 1) {
          return "centos73-x86_64";
        }
      }
    },
    {
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        if (full["Status"] == 0) {
          return "finishInstall";
        } else if (full["Status"] == 1) {
          return "readyForRaid";
        } else if (full["Status"] == 2) {
          return "readyForSys";
        }
      }
    },
    {"mData": 'CreateTime', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<a><i class="glyphicon glyphicon-trash" title="删除" onclick="DeleteOne(\'{{ urlfor "OsInstallController.Cancel" }}\',\'' + full['Id'] + '\')"></i></a>';
      }
    }
    ],
    "aoColumnDefs": [],
    "aaSorting": [[1, "asc"]],
    "fnServerParams": function( aoData )
    {
      var ip = $("#InputIp").val();
      var osversion = $("#selOsVersion").val();
      var status = $("#selStatus").val();
      aoData.push(
        {"name":"filter_ip","value":ip},
        {"name":"filter_osversion","value":osversion},
        {"name":"filter_status","value":status}
        );
    },
    "fnDrawCallback": function() {
      reloadfunc();
    },
    "fnHeaderCallback": function( nHead, aData, iStart, iEnd, aiDisplay ) {

    }
  });

    $(".alert").delay(4000).slideUp(200, function() {
      $(this).alert('close');
    });

    $("#myTable input[name='checkall']").click(function() {
      $("input[name='checksingle']").prop("checked", this.checked);
    });

    $("#myModal").on("hidden.bs.modal", function() {
      $(this).removeData("bs.modal");
    });

  });
</script>
{{end}}