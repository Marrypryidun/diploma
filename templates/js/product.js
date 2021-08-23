"use strict";
(function () {
    $('.form').submit(function(e) {
        e.preventDefault(); // avoid to execute the actual submit of the form.

        var form = $(this);
        var url = form.attr('action');
        var productName=$("#searchProduct").val()
        var searchObj = {"product":productName}
        $.ajax({
            url: url,
            type: 'POST',
            data: searchObj,
            timeout: 15000,
            success: function(result) {
                if(result.products.length==0){
                    errorSearch();
                    return
                }
               // $('#products').tplAppend("item-product", null)
                console.log(result)
            },
            error: function(err) {
                console.log(err);
                errorSearch();
            },
        });
    });
}());

function errorSearch(){
    $('#errorSearch').show()
    setTimeout(() => $('#errorSearch').hide(), 3000);
}