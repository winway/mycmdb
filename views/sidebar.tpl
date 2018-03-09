    <nav id="sidebar">
      <div class="sidebar-header">
        <h3>运维自动化平台</h3>
        <strong><img src="/static/img/favicon.ico" /></strong>
      </div>

      <ul class="list-unstyled components">
        {{ if eq .isadmin 1}}
        <li>
          <a href="{{ urlfor "IdcController.IdcIndexPage" }}">
            <i class="glyphicon glyphicon-dashboard"></i>
            Dashbord
          </a>
        </li>


        <li>
          <a id="SourceManage" href="#SourceManageSubmenu" data-toggle="collapse" aria-expanded="false">
            <i class="glyphicon glyphicon-folder-open"></i>
            资源管理
          </a>
          <ul class="collapse list-unstyled" id="SourceManageSubmenu">
            <li><a href="{{ urlfor "IdcController.IdcIndexPage" }}">机房管理</a></li>
            <li><a href="{{ urlfor "IpController.IpIndexPage" }}">IP管理</a></li>
            <li><a href="{{ urlfor "ServerController.ServerIndexPage" }}">服务器管理</a></li>
          </ul>
        </li>
        {{ end }}

        {{ if eq .isadmin 1}}
        <li>
          <a href="{{ urlfor "OsInstallController.IndexPage" }}">
            <i class="glyphicon glyphicon-wrench"></i>
            OS安装
          </a>
        </li>
        {{ end }}

        <li>
          <a href="{{ urlfor "ReleaseController.IndexPage" }}">
            <i class="glyphicon glyphicon-share-alt"></i>
            代码发布
          </a>
        </li>

        {{ if eq .isadmin 1}}
        <li>
          <a id="userManage" href="#userManageSubmenu" data-toggle="collapse" aria-expanded="false">
            <i class="glyphicon glyphicon-user"></i>
            用户管理
          </a>
          <ul class="collapse list-unstyled" id="userManageSubmenu">
            <li><a href="#">用户管理</a></li>
            <li><a href="#">组管理</a></li>
            <li><a href="#">角色管理</a></li>
            <li><a href="#">权限管理</a></li>
          </ul>
        </li>
      </ul>
      {{ end }}

      {{ if eq .isadmin 1}}
      <ul class="list-unstyled CTAs">
        <li><a href="#" class="download">监控系统</a></li>
        <li><a href="#" class="article">日志系统</a></li>
      </ul>
      {{ end }}
    </nav>