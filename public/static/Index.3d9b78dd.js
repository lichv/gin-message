import{a as e}from"./index.b490ab0f.js";import{r as l,o as a,c as t,w as o,g as i,F as r,e as s,t as n,f as u,h as d}from"./vendor.a921d300.js";const p={name:"HelloWorld",data:()=>({showProductDistribution:!1,showStatusDistribution:!1,list:[],page:{page:1,size:50,total:1,last:1},query:{page:1,size:50,order:"seq_no",sort:1,is_antiepileptic:0,is_clinic_research:0,is_import:0,is_vip:0,reanalysis:0},form:{},status:[{label:"样本开始检测",value:"2"},{label:"检测完成",value:"3"},{label:"开始生信分析",value:"4"},{label:"生信分析完成",value:"5"},{label:"生成临床报告",value:"6"},{label:"临床报告发送完成",value:"7"},{label:"Sanger验证",value:"8"},{label:"验证完成",value:"9"},{label:"生成验证报告",value:"10"},{label:"核酸提取",value:"11"},{label:"文库构建",value:"12"},{label:"外送测序",value:"13"},{label:"生成原始报告",value:"14"},{label:"生成初步报告",value:"15"},{label:"自建库重测",value:"20"},{label:"外送补测",value:"21"},{label:"外送重测",value:"22"},{label:"完毕",value:"32"},{label:"追加验证结果完成",value:"35"},{label:"报告初审",value:"40"},{label:"报告终审",value:"41"},{label:"验证报告初审",value:"42"},{label:"验证报告终审",value:"43"},{label:"推广审核完成",value:"44"},{label:"推广审核驳回",value:"45"},{label:"销售审核完成",value:"46"},{label:"销售审核驳回",value:"47"},{label:"报告审核驳回",value:"48"},{label:"验证报告审核驳回",value:"49"},{label:"报告待打包",value:"50"},{label:"上机测序",value:"51"},{label:"自测序补测",value:"52"},{label:"医学可出报告",value:"53"},{label:"纸质报告邮寄",value:"54"},{label:"数据已发送",value:"55"},{label:"问题样本",value:"56"},{label:"筛点完成",value:"57"},{label:"异常周期样本",value:"504"}]}),props:{msg:String},watch:{$router:{handler(e,l){console.log("val",e),console.log("oldval",l)}},isCollapse(e,l){this.showCollapse=e}},mounted(){this.getPage(1)},methods:{handleShowProductDistribution(){this.form=this.query,"page"in this.form&&delete this.form.page,"size"in this.form&&delete this.form.size,"order"in this.form&&delete this.form.order,"sort"in this.form&&delete this.form.sort,this.showProductDistribution=!0,this.$refs&&this.$refs.product&&this.$refs.product.getData()},handleShowStatusDistribution(){this.form=this.query,"page"in this.form&&delete this.form.page,"size"in this.form&&delete this.form.size,"order"in this.form&&delete this.form.order,"sort"in this.form&&delete this.form.sort,this.showStatusDistribution=!0,this.$refs&&this.$refs.status&&this.$refs.status.getData()},reset(){this.query={page:1,size:20,is_antiepileptic:!1,is_clinic_research:!1,is_import:!1,is_vip:!1,reanalysis:!1},this.getPage(1)},search(){this.getPage(1)},getPage(l){const a=this;this.query.page=l,console.log(this.query),e.getMonitorPages.send(this.query).then((e=>{console.log("getPage的结果",e),2e3==e.state&&(a.list=e.data,a.page=e.page,a.query.page=a.page.page+1)})).catch((e=>{console.log(e)}))},handleNext(){var e=1;this.page.page<this.page.last&&(e=this.page.page+1),this.getPage(e)},handlePrev(){var e=1;this.page.page>2&&(e=this.page.page-1),this.getPage(e)},handleChangePage(e){this.getPage(e)},handleChangeSize(e){this.query.size=e,this.getPage(1)},headerClick(e,l){console.log("column",e),"seq_no"!=e.property&&"status"!=e.property&&"timely"!=e.property&&"user"!=e.property&&"time"!=e.property||("seq_no"==e.property||"status"==e.property||"timely"==e.property?this.query.order=e.property:"user"==e.property?this.query.order="patient":"time"==e.property&&(this.query.order="initial_dt"),2==this.query.sort?this.query.sort=0:this.query.sort=this.query.sort+1,this.getPage(1))},copy(e){var l=document.createElement("input");l.value=e,document.body.appendChild(l),l.select(),document.execCommand("Copy"),document.body.removeChild(l),this.$message({showClose:!0,message:"复制成功",type:"success"})},handleCopyAS(){const l=this;e.getSeqNos.send(this.query).then((e=>{if(console.log(e),2e3==e.state){var a=e.data;console.log("list",a);var t="'"+a.join("','")+"'";console.log("text",t),l.copy(t)}})).catch((e=>{console.log(e)}))}}},h={class:"page_header"},c=d("所有"),m=d("是"),b=d("否"),_=d("所有"),g=d("是"),f=d("否"),y=d("所有"),v=d("是"),w=d("否"),C=d("所有"),q=d("是"),k=d("否"),V=d("所有"),z=d("是"),x=d("否"),P=d("查询"),S=d("重置"),D=d("产品分布"),U=d("状态分布"),$=d("批量复制AS"),H={class:"page_body"},M={key:0},N={key:1},L={key:0},T={key:1},j={key:2},A=i("p",null,"报告备注",-1),I=i("p",null,"其他备注",-1),B={class:"page_footer"},E={class:"batch-handle"},F=i("div",{class:"selected-title"},[d("已选中 "),i("span",{class:"selected—count"},"0"),d(" 项")],-1),W={class:"vux-flexbox selection-items-box vux-flex-row"},G=d("批量转移"),J=d("批量审核"),K=d("批量删除"),O={class:"pagination"};p.render=function(e,d,p,Q,R,X){const Y=l("el-input"),Z=l("el-form-item"),ee=l("el-option"),le=l("el-select"),ae=l("el-radio-button"),te=l("el-radio-group"),oe=l("el-button"),ie=l("el-form"),re=l("el-table-column"),se=l("el-popover"),ne=l("el-table"),ue=l("el-pagination"),de=l("el-monitor-product-distribution"),pe=l("el-dialog"),he=l("el-monitor-status-distribution"),ce=l("el-card");return a(),t(ce,{class:"monitor_page page_wrapper"},{default:o((()=>[i("div",h,[i(ie,{inline:!0,model:R.query,class:"monitot-form-inline"},{default:o((()=>[i(Z,{label:"产品名称：","label-width":"84px"},{default:o((()=>[i(Y,{modelValue:R.query.seq_type,"onUpdate:modelValue":d[1]||(d[1]=e=>R.query.seq_type=e),placeholder:"产品名称",clearable:""},null,8,["modelValue"])])),_:1}),i(Z,{label:"状态：","label-width":"84px"},{default:o((()=>[i(le,{modelValue:R.query.status,"onUpdate:modelValue":d[2]||(d[2]=e=>R.query.status=e),placeholder:"所有状态",onChange:X.search},{default:o((()=>[i(ee,{label:"所有状态",value:"0"}),(a(!0),t(r,null,s(R.status,(e=>(a(),t(ee,{label:e.label,value:e.value},null,8,["label","value"])))),256))])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"超出时长：","label-width":"84px"},{default:o((()=>[i(le,{modelValue:R.query.timely,"onUpdate:modelValue":d[3]||(d[3]=e=>R.query.timely=e),placeholder:"区域",onChange:X.search},{default:o((()=>[i(ee,{label:"14自然日",value:"336"}),i(ee,{label:"15自然日",value:"360"}),i(ee,{label:"20自然日",value:"480"}),i(ee,{label:"25自然日",value:"600"}),i(ee,{label:"28自然日",value:"672"}),i(ee,{label:"1个月",value:"720"}),i(ee,{label:"1000小时",value:"1000"}),i(ee,{label:"2个月",value:"1440"})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"癫痫1.0：","label-width":"84px"},{default:o((()=>[i(te,{size:"mini",modelValue:R.query.is_antiepileptic,"onUpdate:modelValue":d[4]||(d[4]=e=>R.query.is_antiepileptic=e),onChange:X.search},{default:o((()=>[i(ae,{label:"0"},{default:o((()=>[c])),_:1}),i(ae,{label:"1"},{default:o((()=>[m])),_:1}),i(ae,{label:"-1"},{default:o((()=>[b])),_:1})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"是否科研：","label-width":"84px"},{default:o((()=>[i(te,{size:"mini",modelValue:R.query.is_clinic_research,"onUpdate:modelValue":d[5]||(d[5]=e=>R.query.is_clinic_research=e),onChange:X.search},{default:o((()=>[i(ae,{label:"0"},{default:o((()=>[_])),_:1}),i(ae,{label:"1"},{default:o((()=>[g])),_:1}),i(ae,{label:"-1"},{default:o((()=>[f])),_:1})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"重要标志位：","label-width":"96px"},{default:o((()=>[i(te,{size:"mini",modelValue:R.query.is_import,"onUpdate:modelValue":d[6]||(d[6]=e=>R.query.is_import=e),onChange:X.search},{default:o((()=>[i(ae,{label:"0"},{default:o((()=>[y])),_:1}),i(ae,{label:"1"},{default:o((()=>[v])),_:1}),i(ae,{label:"-1"},{default:o((()=>[w])),_:1})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"是否VIP：","label-width":"96px"},{default:o((()=>[i(te,{size:"mini",modelValue:R.query.is_vip,"onUpdate:modelValue":d[7]||(d[7]=e=>R.query.is_vip=e),onChange:X.search},{default:o((()=>[i(ae,{label:"0"},{default:o((()=>[C])),_:1}),i(ae,{label:"1"},{default:o((()=>[q])),_:1}),i(ae,{label:"-1"},{default:o((()=>[k])),_:1})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,{label:"是否重分析：","label-width":"96px"},{default:o((()=>[i(te,{size:"mini",modelValue:R.query.reanalysis,"onUpdate:modelValue":d[8]||(d[8]=e=>R.query.reanalysis=e),onChange:X.search},{default:o((()=>[i(ae,{label:"0"},{default:o((()=>[V])),_:1}),i(ae,{label:"1"},{default:o((()=>[z])),_:1}),i(ae,{label:"-1"},{default:o((()=>[x])),_:1})])),_:1},8,["modelValue","onChange"])])),_:1}),i(Z,null,{default:o((()=>[i(oe,{type:"primary",size:"mini",onClick:X.search},{default:o((()=>[P])),_:1},8,["onClick"]),i(oe,{size:"mini",onClick:X.reset},{default:o((()=>[S])),_:1},8,["onClick"]),i(oe,{size:"mini",onClick:X.handleShowProductDistribution},{default:o((()=>[D])),_:1},8,["onClick"]),i(oe,{size:"mini",onClick:X.handleShowStatusDistribution},{default:o((()=>[U])),_:1},8,["onClick"]),i(oe,{size:"mini",onClick:X.handleCopyAS},{default:o((()=>[$])),_:1},8,["onClick"])])),_:1})])),_:1},8,["model"])]),i("div",H,[i(ne,{data:R.list,style:{width:"100%"},onHeaderClick:X.headerClick},{default:o((()=>[i(re,{fixed:"",prop:"seq_no",label:"检测编号⧫",width:"80"},{default:o((e=>[i("p",{onClick:l=>X.copy(e.row.seq_no)},n(e.row.seq_no),9,["onClick"])])),_:1}),i(re,{prop:"user",label:"用户⧫",width:"80"},{default:o((e=>[i("p",null,"用户："+n(e.row.name),1),i("p",null,"性别："+n(e.row.gender),1),i("p",{onClick:l=>X.copy(e.row.patient)},n(e.row.patient),9,["onClick"]),e.row.is_proband?(a(),t("p",M,"先证者")):(a(),t("p",N,"亲属")),i("p",{onClick:l=>X.copy(e.row.family_no)},n(e.row.family_no),9,["onClick"])])),_:1}),i(re,{prop:"seq_type",label:"产品名称"},{default:o((e=>[i("p",{onClick:l=>X.copy(e.row.seq_type)},n(e.row.seq_type),9,["onClick"])])),_:1}),i(re,{prop:"seq_company",label:"送检公司",width:"100"},{default:o((e=>[i("p",{onClick:l=>X.copy(e.row.seq_company)},n(e.row.seq_company),9,["onClick"])])),_:1}),i(re,{prop:"urgent_level",label:"普通/加急/首次送检",width:"140"},{default:o((e=>[1==e.row.urgent_level?(a(),t("span",L,"普通")):2==e.row.urgent_level?(a(),t("span",T,"加急")):(a(),t("span",j,"首次送检"))])),_:1}),i(re,{prop:"status",label:"样本状态⧫",width:"120"},{default:o((l=>[i("p",null,n(e.$filters.findStatusByCode(l.row.status,R.status)),1)])),_:1}),i(re,{prop:"lab",label:"实验室",width:"160"},{default:o((e=>[i("p",{onClick:l=>X.copy(e.row.sample_lab_no)},"样本编号："+n(e.row.sample_lab_no),9,["onClick"]),i("p",{onClick:l=>X.copy(e.row.sample_type)},"样本类型："+n(e.row.sample_type),9,["onClick"]),i("p",{onClick:l=>X.copy(e.row.lab_seq_type)},"检测类型："+n(e.row.lab_seq_type),9,["onClick"])])),_:1}),i(re,{prop:"people",label:"负责人",width:"120"},{default:o((e=>[i("p",{onClick:l=>X.copy(e.row.auditor_ad)},"医学顾问："+n(e.row.auditor_ad),9,["onClick"]),i("p",{onClick:l=>X.copy(e.row.sale)},"销售员："+n(e.row.sale),9,["onClick"]),i("p",{onClick:l=>X.copy(e.row.department)},"处理部门："+n(e.row.department),9,["onClick"])])),_:1}),i(re,{prop:"timely",label:"花费时长⧫",width:"80"}),i(re,{prop:"remark",label:"备注",width:"80"},{default:o((e=>[""!=e.row.trial_remark?(a(),t(se,{key:0,placement:"top-start",width:200,trigger:"hover"},{reference:o((()=>[A])),default:o((()=>[i("p",{innerHTML:e.row.trial_remark},null,8,["innerHTML"])])),_:2},1024)):u("",!0),""!=e.row.remark?(a(),t(se,{key:1,placement:"top-start",width:200,trigger:"hover"},{reference:o((()=>[I])),default:o((()=>[i("p",{innerHTML:e.row.remark},null,8,["innerHTML"])])),_:2},1024)):u("",!0)])),_:1}),i(re,{label:"标记",width:"150"},{default:o((e=>[i("p",null,"是否癫痫1.0："+n(e.row.is_antiepileptic),1),i("p",null,"是否科研样本："+n(e.row.is_clinic_research),1),i("p",null,"是否重要标志位："+n(e.row.is_import),1),i("p",null,"是否VIP："+n(e.row.is_vip),1),i("p",null,"是否重分析："+n(e.row.reanalysis),1)])),_:1}),i(re,{prop:"time",label:"时间⧫",width:"180"},{default:o((e=>[i("p",null,"采样："+n(e.row.collection_date),1),i("p",null,"录单："+n(e.row.initial_dt),1),i("p",null,"修改："+n(e.row.latest_time),1)])),_:1})])),_:1},8,["data","onHeaderClick"])]),i("div",B,[i("div",E,[F,i("div",W,[i(oe,{size:"mini",disabled:""},{default:o((()=>[G])),_:1}),i(oe,{size:"mini",disabled:""},{default:o((()=>[J])),_:1}),i(oe,{size:"mini",disabled:""},{default:o((()=>[K])),_:1})])]),i("div",O,[i(ue,{small:"",layout:"total, sizes, prev, pager, next","page-size":R.page.size,total:R.page.total,onNextClick:X.handleNext,onPrevClick:X.handlePrev,onCurrentChange:X.handleChangePage,onSizeChange:X.handleChangeSize},null,8,["page-size","total","onNextClick","onPrevClick","onCurrentChange","onSizeChange"])])]),i(pe,{title:"产品分布",modelValue:R.showProductDistribution,"onUpdate:modelValue":d[9]||(d[9]=e=>R.showProductDistribution=e),width:"1600px"},{default:o((()=>[i(de,{ref:"product",query:R.form},null,8,["query"])])),_:1},8,["modelValue"]),i(pe,{title:"状态分布",modelValue:R.showStatusDistribution,"onUpdate:modelValue":d[10]||(d[10]=e=>R.showStatusDistribution=e),width:"1080px"},{default:o((()=>[i(he,{ref:"status",query:R.form},null,8,["query"])])),_:1},8,["modelValue"])])),_:1})};export default p;
