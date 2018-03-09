{{template "base.tpl" .}}

{{define "head"}}
{{end}}

{{define "body"}}

<div class="page-header">
  <h2>服务器管理&nbsp;&nbsp;&nbsp;&nbsp;<small>服务器{{ .serverCnt }}台 已使用{{ .usedServerCnt }}台</small></h2>
</div>

{{ if .flash.error }}
<div class="alert alert-danger" role="alert">{{ .flash.error }}</div>
{{ else if .flash.notice }}
<div class="alert alert-success" role="alert">{{ .flash.notice }}</div>
{{ end }}

<div>
  <form action="" class="form-inline">
    <div class="form-group">
      <label  class="control-label" for="InputKeyword">关键字:</label>
      <input id="InputKeyword" class="form-control" placeholder="请输入关键字">

      <label  class="control-label" for="selIdc">机房:</label>
      <select id="selIdc" multiple class="selectpicker" title="请选择机房" data-style="btn-default">
        {{range $k, $v := .nameSet}}
        <option value={{ $k }}>{{ $k }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selBrand">品牌:</label>
      <select id="selBrand" multiple class="selectpicker" title="请选择品牌" data-style="btn-default">
        {{range $index, $elem := .brands}}
        <option value={{ $elem.Brand }}>{{ $elem.Brand }}</option>
        {{end}}
      </select>

      <label  class="control-label" for="selOsVersion">OS:</label>
      <select id="selOsVersion" multiple class="selectpicker" title="请选择OS" data-style="btn-default">
        <option value=0>CentOS-6.9-x86_64</option>
        <option value=1>centos73-x86_64</option>
      </select>

      <label  class="control-label" for="selStatus">状态:</label>
      <select id="selStatus" multiple class="selectpicker" title="请选择状态" data-style="btn-default">
        <option value=0>未使用</option>
        <option value=1>已使用</option>
      </select>

      <a onclick="Search()" class="btn btn-primary"><i class="glyphicon glyphicon-search"></i>查询</a>

      <a class="btn btn-primary" href="{{ urlfor "ServerController.ServerAddPage" }}" data-toggle="modal" data-target="#myModal">
        <i class="glyphicon glyphicon-plus"></i>添加
      </a>

      <a onclick="Delete('{{ urlfor "ServerController.DeleteServer" }}')" class="btn btn-primary"><i class="glyphicon glyphicon-trash"></i>删除</a>
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
        <th>SN</th>
        <th>机房</th>
        <th>主机名</th>
        <th>远程卡IP</th>
        <th>网卡IP</th>
        <th>配置</th>
        <th>OS</th>
        <th>品牌</th>
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
    "sAjaxSource": "{{ urlfor "ServerController.ListServer" }}",
    "aoColumns": [
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<input id=' + full['Sn'] + ' value=' + full['Sn'] + ' name="checksingle" type="checkbox">';
      }
    },
    {"mData": 'Sn', "sClass": "alignRight"},
    {"mData": 'Idc.Name', "sClass": "alignRight"},
    {"mData": 'HostName', "sClass": "alignRight"},
    {"mData": 'RemoteCardIp', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return 'Eth1: ' + full['Eth1Ip'] + '<br />' + 'Eth2: ' + full['Eth2Ip'] + '<br />' + 'Eth3: ' + full['Eth3Ip'] + '<br />' + 'Eth4: ' + full['Eth4Ip'];
      }
    },
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return 'Cpu: ' + full['Cpu'] + '<br />' + 'Disk: ' + full['Disk'] + '<br />' + 'Memory: ' + full['Memory'];
      }
    },
    {
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        if (full["OsVersion"] == 0) {
          return "CentOS-6.9-x86_64";
        } else if (full["OsVersion"] == 1) {
          return "centos73-x86_64";
        }
      }
    },
    {"mData": 'Brand', "sClass": "alignRight"},
    {"mData": 'Status', "sClass": "alignRight"},
    {"mData": 'CreateTime', "sClass": "alignRight"},
    {
      "bSortable": false,
      "sClass": "alignRight",
      "mRender": function (data, type, full) {
        return '<a href="{{ urlfor "ServerController.ServerDetailPage" ":id" "" }}' + full['Sn'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-info-sign" title="详情"></i></a> | <a href="{{ urlfor "ServerController.ServerEditPage" ":id" "" }}' + full['Sn'] + '" data-toggle="modal" data-target="#myModal"><i class="glyphicon glyphicon-edit" title="编辑"></i></a> | <a><i class="glyphicon glyphicon-trash" title="删除" onclick="DeleteOne(\'{{ urlfor "ServerController.DeleteServer" }}\',\'' + full['Sn'] + '\')"></i></a> | <a><i class="glyphicon glyphicon-eye-close" title="密码" onclick=""></i></a>';
      }
    }
    ],
    "aoColumnDefs": [],
    "aaSorting": [[1, "asc"]],
    "fnServerParams": function( aoData )
    {
      var keyword = $("#InputKeyword").val();
      var idc = $("#selIdc").val();
      var brand = $("#selBrand").val();
      var osversion = $("#selOsVersion").val();
      var status = $("#selStatus").val();
      aoData.push(
        {"name":"filter_keyword","value":keyword},
        {"name":"filter_idc","value":idc},
        {"name":"filter_brand","value":brand},
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