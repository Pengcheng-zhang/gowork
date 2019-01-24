//Loads the correct sidebar on window load,
//collapses the sidebar on window resize.
// Sets the min-height of #page-wrapper to window size
$(function () {
	$('#side-menu').metisMenu();
	$(window).bind("load resize", function () {
		var topOffset = 50;
		var width = (this.window.innerWidth > 0) ? this.window.innerWidth : this.screen.width;
		if (width < 768) {
			$('div.navbar-collapse').addClass('collapse');
			topOffset = 100; // 2-row-menu
		} else {
			$('div.navbar-collapse').removeClass('collapse');
		}

		var height = ((this.window.innerHeight > 0) ? this.window.innerHeight : this.screen.height) - 1;
		height = height - topOffset;
		if (height < 1) height = 1;
		if (height > topOffset) {
			$("#page-wrapper").css("min-height", (height) + "px");
		}
	});
});
var num = 0, oUl = $("#min_title_list"), hide_nav = $("#Hui-tabNav");
var globalEditor = null;
var tableLocalOptions = {
	"sProcessing": "处理中...",
	"sLengthMenu": "显示 _MENU_ 项结果",
	"sZeroRecords": "没有匹配结果",
	"sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
	"sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
	"sInfoFiltered": "(由 _MAX_ 项结果过滤)",
	"sInfoPostFix": "",
	"sSearch": "搜索:",
	"sUrl": "",
	"sEmptyTable": "表中数据为空",
	"sLoadingRecords": "载入中...",
	"sInfoThousands": ",",
	"oPaginate": {
		"sFirst": "首页",
		"sPrevious": "上页",
		"sNext": "下页",
		"sLast": "末页"
	},
	"oAria": {
		"sSortAscending": ": 以升序排列此列",
		"sSortDescending": ": 以降序排列此列"
	}
 };
 var articleTable = null;
/*获取顶部选项卡总长度*/
function tabNavallwidth() {
	var taballwidth = 0,
		$tabNav = hide_nav.find(".acrossTab"),
		$tabNavWp = hide_nav.find(".Hui-tabNav-wp"),
		$tabNavitem = hide_nav.find(".acrossTab li"),
		$tabNavmore = hide_nav.find(".Hui-tabNav-more");
	if (!$tabNav[0]) { return }
	$tabNavitem.each(function (index, element) {
		taballwidth += Number(parseFloat($(this).width() + 60))
	});
	$tabNav.width(taballwidth + 25);
	var w = $tabNavWp.width();
	if (taballwidth + 25 > w) {
		$tabNavmore.show()
	}
	else {
		$tabNavmore.hide();
		$tabNav.css({ left: 0 });
	}
}

/*最新tab标题栏列表*/
function min_titleList() {
	var topWindow = $(window.parent.document),
		show_nav = topWindow.find("#min_title_list"),
		aLi = show_nav.find("li");
}

/*创建iframe*/
function creatIframe(href, titleName) {
	var topWindow = $(window.parent.document),
		show_nav = topWindow.find('#min_title_list'),
		iframe_box = topWindow.find('#iframe_box'),
		iframeBox = iframe_box.find('.show_iframe'),
		$tabNav = topWindow.find(".acrossTab"),
		$tabNavWp = topWindow.find(".Hui-tabNav-wp"),
		$tabNavmore = topWindow.find(".Hui-tabNav-more");
	var taballwidth = 0;

	show_nav.find('li').removeClass("active");
	show_nav.append('<li class="active" id="nav_'+href+'"><span data-href="' + href + '">' + titleName + '</span><i></i><em></em></li>');

	var $tabNavitem = topWindow.find(".acrossTab li");
	if (!$tabNav[0]) { return }
	$tabNavitem.each(function (index, element) {
		taballwidth += Number(parseFloat($(this).width() + 60))
	});
	$tabNav.width(taballwidth + 25);
	var w = $tabNavWp.width();
	if (taballwidth + 25 > w) {
		$tabNavmore.show()
	}
	else {
		$tabNavmore.hide();
		$tabNav.css({ left: 0 })
	}
	iframeBox.hide();
	var tpl = template($('#' + href).html());
	var html = tpl(1234);
	iframe_box.append(html);
	if (href == "editArticle") {
		initEditor();
		getArticleTypeList();
	}else if(href == "articleList") {
		initArticleList();
	}
}

function initEditor() {
	var E = window.wangEditor
	globalEditor = new E('#article_content')
	globalEditor.create()
	globalEditor.txt.html('')
}

function initArticleList() {
	articleTable = $('#table_article_list').DataTable({
		responsive: true,
		language: tableLocalOptions,
		"serverSide": true,
		"bFilter": false,
		"paging": true,
		"columns": [
			{"data": "id"},
			{"data": "title"},
			{"data": "created_at"},
			{"data": "status"},
			{
				"data": "view_times"
			}
		],
		"columnDefs": [
			{
				"orderable": false,
				"targets": [0, 1, 2, 3, 4, 5]
			},
			{
				"targets": 0,
				"render": function(data, type, row, meta){
					return '<a href="/article/'+data+'" target="_blank">'+data+'</a>';
				}
			},
			{
				"targets": 3,
				"render": function(data, type, row, meta){
					if (data == "B") {
						return "被打回";
					}else{
						return "正常";
					}
				}
			},
			{
				"targets": 5,
				"render": function(data, type, row, meta){
					return '<a href="JavaScript:updateArticle('+row['id']+')"><i class="fa fa-edit"></i></a>';
				}
			}
		],
		"ajax": {
			"url": "http://localhost:3000/admin/article_list",
			"type": "POST",
			"data" : function(d) {
				 d.status = $('#article_status').val();
				 d.keyword = $('#article_keyword').val();  
			}
		},

	});
}

function getArticleTypeList() {
	$.post("/article/api_type_list", {}, function(data) {
		if(data.code == 10000) {
			for(var i=0;i<data.result.length;i++) {
				$('#article_type').append('<option value="'+data.result[i].id+'">'+data.result[i].name+'</option>');
			}
		}else {

		}
	});
}

function editArticle() {
	var articleData = new Object;
	var id = $('#article_id').val();
	if (id > 0) {
		articleData.id = id;
	}
	articleData.title = $('#article_title').val();
	articleData.type = $('#article_type').val();
	articleData.content = globalEditor.txt.html();
	var errMsg = ''
	if(articleData.title == '') {
		errMsg = "请输入标题"
	}else if(articleData.type <= 0) {
		errMsg = "请选择类型"
	}else if(articleData.content.length < 20) {
		errMsg = "请输入不少于20字的内容"
	}
	if (errMsg.length > 0) {
		$.toast({
			text : errMsg,
			position: 'mid-center',
			showHideTransition: 'fade',
			icon: 'error',
			loader: false, 
		});
		return;
	}
	var req_url = "/article/api_new";
	if (id > 0) {
		req_url = "/article/api_update";
	}
	$.post(req_url, articleData, function(data) {
		if(data.code == 10000) {
			$.toast({
				text : "发布成功",
				position: 'mid-center',
				showHideTransition: 'fade',
				icon: 'success',
				loader: false, 
			});
		}else {
			$.toast({
				text : data.message,
				position: 'mid-center',
				showHideTransition: 'fade',
				icon: 'error',
				loader: false, 
			});
		}
	});
}
function updateArticle(id) {
	$('#side_nav_edit_article').trigger('click');
	$.post("/article/api_detail", {id: id}, function(data) {
		if (data.code == 10000) {
			$('#article_id').val(data.result.id);
			$('#article_title').val(data.result.title);
			$('#article_type').val(data.result.type);
			globalEditor.txt.html(data.result.content);
		}
	});
}
function reloadArticleTable() {
	articleTable.draw(false);
}

/*时间*/
function getHTMLDate(obj) {
	var d = new Date();
	var weekday = new Array(7);
	var _mm = "";
	var _dd = "";
	var _ww = "";
	weekday[0] = "星期日";
	weekday[1] = "星期一";
	weekday[2] = "星期二";
	weekday[3] = "星期三";
	weekday[4] = "星期四";
	weekday[5] = "星期五";
	weekday[6] = "星期六";
	_yy = d.getFullYear();
	_mm = d.getMonth() + 1;
	_dd = d.getDate();
	_ww = weekday[d.getDay()];
	obj.html(_yy + "年" + _mm + "月" + _dd + "日 " + _ww);
};

$(function () {
	getHTMLDate($("#top_time"));
	var resizeID;

	/*选项卡导航*/
	$(".sidebar-nav").on("click", ".nav a", function () {
		var href = $(this).data('href');
		if(!href||href==""){
			return false;
		}
		if($('#nav_'+href).length > 0) {
			$("#min_title_list li").removeClass('active');
			$('#nav_'+href).addClass('active');
			$('#iframe_box .show_iframe').hide();
			$('#show_admin_'+href).show();
		}else{
			creatIframe(href, $(this).data('title'));
		}
	});

	$(document).on("click","#min_title_list li",function(){
		var href = $(this).attr('id').split('_')[1];
		if(!$(this).hasClass('active')) {
			$("#min_title_list li").removeClass('active');
			$(this).addClass('active');
			$('#iframe_box .show_iframe').hide();
			$('#show_admin_'+href).show();
		}
	});
	$(document).on("click", "#min_title_list li i", function () {
		var index = $(this).parent("li").index();
		var frameName = $(this).parents("li").find('span').data('href');
		$(this).parent('li').remove();
		$('#show_admin_' + frameName).remove();
		$("#min_title_list li").eq(index-1).trigger("click");
	});

	$(document).on("click", "#query_article", function() {
		reloadArticleTable();
	});

	$(document).on("click", "#edit_article", function(){
		editArticle();
	});
	function toNavPos() {
		oUl.stop().animate({ 'left': -num * 100 }, 100);
	}
}); 
