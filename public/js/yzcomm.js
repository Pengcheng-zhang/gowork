function regist(email, password) {
    $.post("/api/login",{"email":email, "password":password}, function(data) {
        if(data.result) {
            window.location.href = data.next_url;
        }else{
            console.log(data.message);
        }
    });
}