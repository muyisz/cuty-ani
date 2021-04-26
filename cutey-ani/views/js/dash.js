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
                alert("失败1")
            }
        },
        error: function () {
            alert("失败")
        }
    })
})