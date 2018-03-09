    <div class="modal-header">
      <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
        &times;
      </button>
      <h4 class="modal-title" id="myModalLabel">
        服务器详情
      </h4>
    </div>
    <form id="app" class="form-horizontal" method="post" action="">
      <div class="modal-body">
        <div class="form-group">
          <label for="Sn" class="col-sm-3 control-label">*SN：</label>
          <div class="col-sm-6">
            <input v-model="Sn" type="text" class="form-control" name="Sn" placeholder="请输入序列号" readonly>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-3 control-label" for="Idc">*机房：</label>
          <div class="col-sm-6">
            <select v-model="Idc" name="Idc" class="form-control" title="请选择机房" data-style="btn-default" readonly>
              {{range $index, $elem := .names}}
              <option value={{ $elem.Id }}>{{ $elem.Name }}</option>
              {{end}}
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="CabinetNo" class="col-sm-3 control-label">*机柜号：</label>
          <div class="col-sm-6">
            <input v-model="CabinetNo" type="text" class="form-control" name="CabinetNo" placeholder="请输入机柜号" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="IdInsideCabinet" class="col-sm-3 control-label">*机柜内序号：</label>
          <div class="col-sm-6">
            <input v-model="IdInsideCabinet" type="text" class="form-control" name="IdInsideCabinet" placeholder="请输入机柜内序号" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="RemoteCardIp" class="col-sm-3 control-label">*远程卡IP：</label>
          <div class="col-sm-6">
            <input v-model="RemoteCardIp" type="text" class="form-control" name="RemoteCardIp" placeholder="请输入远程卡IP" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="RemoteCardMac" class="col-sm-3 control-label">远程卡Mac：</label>
          <div class="col-sm-6">
            <input v-model="RemoteCardMac" type="text" class="form-control" name="RemoteCardMac" placeholder="请输入远程卡Mac" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="HostName" class="col-sm-3 control-label">主机名：</label>
          <div class="col-sm-6">
            <input v-model="HostName" type="text" class="form-control" name="HostName" placeholder="请输入主机名" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth1Ip" class="col-sm-3 control-label">Eth1 IP：</label>
          <div class="col-sm-6">
            <input v-model="Eth1Ip" type="text" class="form-control" name="Eth1Ip" placeholder="请输入网卡IP" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth2Ip" class="col-sm-3 control-label">Eth2 IP：</label>
          <div class="col-sm-6">
            <input v-model="Eth2Ip" type="text" class="form-control" name="Eth2Ip" placeholder="请输入网卡IP" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth3Ip" class="col-sm-3 control-label">Eth3 IP：</label>
          <div class="col-sm-6">
            <input v-model="Eth3Ip" type="text" class="form-control" name="Eth3Ip" placeholder="请输入网卡IP" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth4Ip" class="col-sm-3 control-label">Eth4 IP：</label>
          <div class="col-sm-6">
            <input v-model="Eth4Ip" type="text" class="form-control" name="Eth4Ip" placeholder="请输入网卡IP" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth1Mac" class="col-sm-3 control-label">Eth1 MAC：</label>
          <div class="col-sm-6">
            <input v-model="Eth1Mac" type="text" class="form-control" name="Eth1Mac" placeholder="请输入网卡Mac" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth2Mac" class="col-sm-3 control-label">Eth2 MAC：</label>
          <div class="col-sm-6">
            <input v-model="Eth2Mac" type="text" class="form-control" name="Eth2Mac" placeholder="请输入网卡Mac" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth3Mac" class="col-sm-3 control-label">Eth3 MAC：</label>
          <div class="col-sm-6">
            <input v-model="Eth3Mac" type="text" class="form-control" name="Eth3Mac" placeholder="请输入网卡Mac" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Eth4Mac" class="col-sm-3 control-label">Eth4 MAC：</label>
          <div class="col-sm-6">
            <input v-model="Eth4Mac" type="text" class="form-control" name="Eth4Mac" placeholder="请输入网卡Mac" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Nic" class="col-sm-3 control-label">网卡：</label>
          <div class="col-sm-6">
            <input v-model="Nic" type="text" class="form-control" name="Nic" placeholder="请输入CPU信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="NicDetail" class="col-sm-3 control-label">网卡详情：</label>
          <div class="col-sm-6">
            <input v-model="NicDetail" type="text" class="form-control" name="NicDetail" placeholder="请输入CPU信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Cpu" class="col-sm-3 control-label">CPU：</label>
          <div class="col-sm-6">
            <input v-model="Cpu" type="text" class="form-control" name="Cpu" placeholder="请输入CPU信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="CpuDetail" class="col-sm-3 control-label">CPU详情：</label>
          <div class="col-sm-6">
            <input v-model="CpuDetail" type="text" class="form-control" name="CpuDetail" placeholder="请输入CPU信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Disk" class="col-sm-3 control-label">硬盘：</label>
          <div class="col-sm-6">
            <input v-model="Disk" type="text" class="form-control" name="Disk" placeholder="请输入硬盘信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="DiskDetail" class="col-sm-3 control-label">硬盘详情：</label>
          <div class="col-sm-6">
            <input v-model="DiskDetail" type="text" class="form-control" name="DiskDetail" placeholder="请输入硬盘信息" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="Memory" class="col-sm-3 control-label">内存：</label>
          <div class="col-sm-6">
            <input v-model="Memory" type="text" class="form-control" name="Memory" placeholder="请输入内存" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="MemoryDetail" class="col-sm-3 control-label">内存详情：</label>
          <div class="col-sm-6">
            <input v-model="MemoryDetail" type="text" class="form-control" name="MemoryDetail" placeholder="请输入内存" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="PowerSupply" class="col-sm-3 control-label">电源：</label>
          <div class="col-sm-6">
            <input v-model="PowerSupply" type="text" class="form-control" name="PowerSupply" placeholder="请输入内存" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="PowerSupplyDetail" class="col-sm-3 control-label">电源详情：</label>
          <div class="col-sm-6">
            <input v-model="PowerSupplyDetail" type="text" class="form-control" name="PowerSupplyDetail" placeholder="请输入内存" readonly>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-3 control-label" for="Brand">品牌：</label>
          <div class="col-sm-6">
            <select v-model="Brand" name="Brand" class="form-control" title="请选择品牌" data-style="btn-default" readonly>
              {{range $index, $elem := .names}}
              <option value={{ $elem.Id }}>{{ $elem.Name }}</option>
              {{end}}
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="WarrantyTime" class="col-sm-3 control-label">过保时间：</label>
          <div class="col-sm-6">
            <input v-model="WarrantyTime" type="text" class="form-control" name="WarrantyTime" placeholder="请输入过保时间" readonly>
          </div>
        </div>

        <div class="form-group">
          <label for="RaidInfo" class="col-sm-3 control-label">RAID信息：</label>
          <div class="col-sm-6">
            <input v-model="RaidInfo" type="text" class="form-control" name="RaidInfo" placeholder="请输入内存" readonly>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-3 control-label" for="OsVersion">操作系统：</label>
          <div class="col-sm-6">
            <select v-model="OsVersion" name="OsVersion" class="form-control" title="请选择操作系统" data-style="btn-default" readonly>
              <option value=0>CentOS-6.9-x86_64</option>
              <option value=1>centos73-x86_64</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label  class="col-sm-3 control-label" for="Status">状态：</label>
          <div class="col-sm-6">
            <select v-model="Status" name="Status" class="form-control" title="请选择状态" data-style="btn-default" readonly>
              <option value=0>未使用</option>
              <option value=1>已使用</option>
            </select>
          </div>
        </div>

        <div class="form-group">
          <label for="Comment" class="col-sm-3 control-label">备注：</label>
          <div class="col-sm-6">
            <input v-model="Comment" type="text" class="form-control" name="Comment" placeholder="备注" readonly>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="button" class="btn btn-primary" disabled>保存</button>
      </div>
    </form>

    <script type="text/javascript">
      var vue = new Vue({
        el: '#app',
        data: {
          Sn: {{ .server.Sn }},
          Idc: {{ .server.Idc.Id }},
          CabinetNo: {{ .server.CabinetNo }},
          IdInsideCabinet: {{ .server.IdInsideCabinet }},
          RemoteCardIp: {{ .server.RemoteCardIp }},
          RemoteCardMac: {{ .server.RemoteCardMac }},
          HostName: {{ .server.HostName }},
          Eth1Ip: {{ .server.Eth1Ip }},
          Eth2Ip: {{ .server.Eth2Ip }},
          Eth3Ip: {{ .server.Eth3Ip }},
          Eth4Ip: {{ .server.Eth4Ip }},
          Eth1Mac: {{ .server.Eth1Mac }},
          Eth2Mac: {{ .server.Eth2Mac }},
          Eth3Mac: {{ .server.Eth3Mac }},
          Eth4Mac: {{ .server.Eth4Mac }},
          Nic: {{ .server.Nic }},
          NicDetail: {{ .server.NicDetail }},
          Cpu: {{ .server.Cpu }},
          CpuDetail: {{ .server.CpuDetail }},
          Disk: {{ .server.Disk }},
          DiskDetail: {{ .server.DiskDetail }},
          Memory: {{ .server.Memory }},
          MemoryDetail: {{ .server.MemoryDetail }},
          PowerSupply: {{ .server.PowerSupply }},
          PowerSupplyDetail: {{ .server.PowerSupplyDetail }},
          Brand: {{ .server.Brand }},
          WarrantyTime: {{ .server.WarrantyTime }},
          RaidInfo: {{ .server.RaidInfo }},
          OsVersion: {{ .server.OsVersion }},
          Status: {{ .server.Status }},
          Comment: {{ .server.Comment }}
        }
      });
    </script>