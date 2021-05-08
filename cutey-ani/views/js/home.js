$(document).ready(function () {
    $.ajax({
        url:'/getphoto',
        type:'get',
        dataType:'json',
        processData: false,
        contentType: false,
        success: function(data){
            if(data.pass){
                const k="img"
                for(i=0;i<11;i++){
                    n=i+1
                    document.getElementById(k+n).src=data.url[i]
                }
            }
        },
        error: function(){
            alert("图片获取失败!")
        }
    })
    $('#login').click(function () {
        window.location.href = "/login";
    })
    $('#adopt').click(function(){
        alert("敬请期待")
    })

})