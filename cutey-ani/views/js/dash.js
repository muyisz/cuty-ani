$(document).ready(function () {
    let formData = new FormData();
    $.ajax({
        url: '/info',
        type: 'get',
        dataType: 'json',
        processData: false,
        contentType: false,
        success: function (data) {
            if (data.pass) {
                document.getElementById("setphone").innerText = data.phone
                document.getElementById("setname").innerText = data.nickname
                document.getElementById("setaddress").innerText = data.address
            }else{
                alert("失败")
            }
        },
        error: function () {
            alert("失败")
        }
    })
    $('#up').click(function(){
        const up_photo = $('#photo')[0].files[0];
        var formdata=new FormData();
        formdata.append("photo",up_photo);
        $.ajax({
            url:'/up_photo',
            type:'post',
            async: false,
            cache: false,
            dataType:'json',
            data:formdata,
            processData: false,
            contentType: false,
            success:function(data){
                if(data.pass){
                    alert("上传成功")
                }else{
                    alert("上传失败")
                }
            },
            error:function(){
                alert("上传失败")
            }
        })
    })
})