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
            } else {
                alert("失败")
            }
        },
        error: function () {
            alert("失败")
        }
    })
    $('#up').click(function () {
        const up_photo = $('#photo')[0].files[0];
        if (up_photo) {
            var photo_name = document.getElementById('photo').value;
            var index = photo_name.lastIndexOf('.');
            var ext = photo_name.substring(index);
            if (ext == '.png' || ext == '.jpg' || ext == '.jepg' || ext == '.jfif') {
                var formdata = new FormData();
                formdata.append("photo", up_photo);
                formdata.append("ext",ext);
                $.ajax({
                    url: '/up_photo',
                    type: 'post',
                    async: false,
                    cache: false,
                    dataType: 'json',
                    data: formdata,
                    processData: false,
                    contentType: false,
                    success: function (data) {
                        if (data.pass) {
                            alert("上传成功" + ext)
                        } else {
                            alert("上传失败")
                        }
                    },
                    error: function () {
                        alert("上传失败")
                    }
                })
            }
            else{
                alert("仅仅支持png,jpg,jepg,jfif格式")
            }
        }
        else {
            alert("上传失败")
        }
    })
})