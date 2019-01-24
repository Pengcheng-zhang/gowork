// refresh money
function refreshMoney() {
    $.post('/ajax/money', function (data) {
        $('#money').html(data);
    });
}

var moveEnd = function (obj) {
    obj.focus();
    obj = obj.get(0);
    var len = obj.value.length;
    if (document.selection) {
        var sel = obj.createTextRange();
        sel.moveStart('character', len);
        sel.collapse();
        sel.select();
    } else if (typeof obj.selectionStart == 'number' && typeof obj.selectionEnd == 'number') {
        obj.selectionStart = obj.selectionEnd = len;
    }
}

function dispatch() {
    var q = document.getElementById("q");
    if (q.value != "") {
        var url = 'https://www.google.com/search?q=site:v2ex.com/t%20' + q.value;
        if (navigator.userAgent.indexOf('iPad') > -1 || navigator.userAgent.indexOf('iPod') > -1 || navigator.userAgent.indexOf('iPhone') > -1) {
            location.href = url;
        } else {
            window.open(url, "_blank");
        }
        return false;
    } else {
        return false;
    }
}

function resendVerificationEmail(once) {
    $("#ButtonResendVerification").prop("disabled", true);
    $.post('/settings/resend?once=' + once, function (data) {
        $("#ResendResponse").html(data.message);
    });
}

function goTop() {
    event.preventDefault();
    $("html, body").animate({ scrollTop: 0 }, 1000);
}
function signIn() {
    var username = $('input[name=user_name]').val();
    var password = $('input[name=password]').val();
    if (username.length == 0) {
        showErrorMessage('用户名不能为空');
        return
    }else if(password.length == 0){
        showErrorMessage('请输入正确的密码');
        return
    }
    $.post("/api_login", {email: username, password: password}, function(data) {
        if(data.code == "10000") {
            window.location.href = data.result;
        }else{
            showErrorMessage(data.message);
        }
    });
}
function signUp() {
    var username = $('input[name=username]').val();
    var password = $('input[name=password]').val();
    var email = $('input[name=email]').val();
    if (username.length == 0) {
        showErrorMessage('用户名不能为空');
        return
    }else if(password.length == 0){
        showErrorMessage('请输入正确的密码');
        return
    }else if(email.length == 0 ) {
        showErrorMessage('请输入正确的邮箱');
        return
    }
    if(! checkEmail(email)) {
        showErrorMessage('请输入正确的邮箱');
        return
    }
    $.post("/api_signup", {username: username, password: password, email: email}, function(data) {
        if(data.code == "10000") {
            window.location.href = data.result;
        }else{
            showErrorMessage(data.message);
        }
    });
}
function checkEmail(email) {
    var reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$");
    if(reg.test(email)) {
        return true;
    }
    return false;
}
function showErrorMessage(message) {
    $('#yz_tip_message').removeClass('yz-hidden').addClass('yz-error').append('<i class="fa fa-warning fa-warning-2x"></i>' + message);
    setTimeout(function(){
        $('#yz_tip_message').addClass('yz-hidden').removeClass('yz-error').html('');
    }, 2000);
}

function showSuccessMessage(message) {
    $('#yz_tip_message').removeClass('yz-hidden').addClass('yz-success').append('<i class="fa fa-check-circle fa-check-circle-2x"></i>' + message);
    setTimeout(function(){
        $('#yz_tip_message').addClass('yz-hidden').removeClass('yz-success').html('');
    }, 2000);
}

function signOut() {
    if (confirm('确定要退出登录？')) { 
        $.post('/api_logout', {},  function(data){
            if(data.Code == "10000") {
                window.location.href="/";
            }
        });
    }
}
// reply a reply
function replyOne(username) {
    replyContent = $("#reply_content");
    oldContent = replyContent.val();
    prefix = "@" + username + " ";
    newContent = ''
    if (oldContent.length > 0) {
        if (oldContent != prefix) {
            newContent = oldContent + "\n" + prefix;
        }
    } else {
        newContent = prefix
    }
    replyContent.focus();
    replyContent.val(newContent);
    moveEnd($("#reply_content"));
}

// send a thank to reply
function thankReply(replyId, token) {
    $.post('/thank/reply/' + replyId + "?t=" + token, function () {
        $('#thank_area_' + replyId).addClass("thanked").html("感谢已发送");
        refreshMoney();
    });
}

// send a thank to topic
function thankTopic(topicId, token) {
    $.post('/thank/topic/' + topicId + "?t=" + token, function (data) {
        $('#topic_thank').html('<span class="f11 gray" style="text-shadow: 0px 1px 0px #fff;">感谢已发送</span>');
        refreshMoney();
    });
}

function upVoteTopic(topicId) {
    if (csrfToken) {
        var request = $.ajax({
            url: '/article/api_prise',
            data: { id: topicId },
            type: "POST",
            dataType: "json"
        });
        request.done(function (data) {
            if (data.code == '10000') {
                $('#topic_votes>a:first>li').html(data.result);
                $.toast({
                    text : '感谢点赞',
                    position: 'mid-center',
                    showHideTransition: 'fade',
                    icon: 'success',
                    loader: false, 
                });
            }else{
                $.toast({
                    text : data.message,
                    position: 'mid-center',
                    showHideTransition: 'fade',
                    icon: 'warning',
                    loader: false, 
                });
            }
        });
    }
}

function downVoteTopic(topicId) {
    if (csrfToken) {
        var request = $.ajax({
            url: '/article/api_diss',
            data: { id: topicId },
            type: "POST",
            dataType: "json"
        });
        request.done(function (data) {
            if (data.code == '10000') {
                $('#topic_votes>a:last>li').html(data.result);
            }else{
                $.toast({
                    text : data.message,
                    position: 'mid-center',
                    showHideTransition: 'fade',
                    icon: 'warning',
                    loader: false, 
                });
            }
        });
    }
}
function replyArticle(topicId) {
    var content = $('#yz_reply_content').val();
    if (content.length == 0) {
        showErrorMessage("请输入评论内容！");
        return
    }
    $.post('/article/api_add_comment', {id: topicId, content: content}, function(data) {
        if (data.code == '10000') {
            var tpl = template($('#yz_topic_reply').html());
            var html = tpl(data);
            $('#yz_reply_content').val('');
            $('#yz_reply_list').append(html);
        }else{
            showErrorMessage(data.message);
            return;
        }
    });
}

function forgetPassword() {
    var email = $('input[name=email]').val();
    if(! checkEmail(email)) {
        showErrorMessage('请输入正确的邮箱');
        return
    }
    $.post('/api_forget_password', {email: email}, function(data) {
        if (data.Code == '10000') {
            $.toast({
                text : data.Message,
                position: 'mid-center',
                showHideTransition: 'fade',
                icon: 'success',
                loader: false, 
            });
        }else{
            $.toast({
                text : data.Message,
                position: 'mid-center',
                showHideTransition: 'fade',
                icon: 'error',
                loader: false, 
            });
        }
    });
}

function publishTopic() {
    var errors = 0;
    var em = $("#yz_error_message");
    var node = $('#yz_nodes').val();
    var content = editor.txt.html();

    var title = $("#topic_title").val();

    if (title.length == 0) {
        errors++;
        em.html("主题标题不能为空");
    } else if (title.length > 120) {
        errors++;
        em.html("主题标题不能超过 120 个字符");
    }

    if (content.length > 20000) {
        errors++;
        em.html("主题内容不能超过 20000 个字符");
    }

    if (node <= 0) {
        errors++;
        em.html("请选择一个节点");
    }
    if (errors == 0) {
        $.post("/article/api_new", { "title": title, "content": content, "node":  node}, function (data) {
            if (data.Code == '10000') {
                em.html('发表成功,即将跳转^^');
                setTimeout(function(){
                    location.href = '/article/' + data.result;
                }, 500);
            } else {
                em.html(data.message);
            }
        });
    }
}

function previewTopicSupplement() {
    var box = $("#box");
    var preview = $("#topic_preview");
    if (preview.length == 0) {
        box.append('<div class="dock_area"><div class="inner"><span class="gray">主题附言预览</span></div></div><div class="inner" id="topic_preview"></div>');
        preview = $("#topic_preview");
    }
    var txt = $("#topic_supplement").val();
    var syntax = $("#syntax").val();
    if (syntax == 0) {
        $.post("/preview/default", { 'txt': txt }, function (data) {
            preview.html('<div class="topic_content"><div class="markdown_body">' + data + '</div></div>');
        });
    }
    if (syntax == 1) {
        $.post("/preview/markdown", { 'md': txt }, function (data) {
            preview.html('<div class="topic_content"><div class="markdown_body">' + data + '</div></div>');
        });
    }
}

function previewTopicContent() {
    var box = $("#box");
    var preview = $("#topic_preview");
    var syntax = $("#syntax").val();
    var syntax_text = $('#syntax option:selected').text()
    if (preview.length == 0) {
        box.append('<div class="dock_area"><div class="inner"><div class="fr gray" id="syntax_text">文本标记语法 &nbsp;<strong>' + syntax_text + '</strong></div><span class="gray">主题内容预览</span></div></div><div class="inner" id="topic_preview"></div>');
        preview = $("#topic_preview");
    } else {
        $("#syntax_text").html("文本标记语法 &nbsp;<strong>" + syntax_text + "</strong>");
    }
    var txt = $("#topic_content").val();
    if (syntax == 0) {
        $.post("/preview/default", { 'txt': txt }, function (data) {
            preview.html('<div class="topic_content"><div class="markdown_body">' + data + '</div></div>');
        });
    }
    if (syntax == 1) {
        $.post("/preview/markdown", { 'md': txt }, function (data) {
            preview.html('<div class="topic_content"><div class="markdown_body">' + data + '</div></div>');
        });
    }
}

// End: draft management for /notes/new
