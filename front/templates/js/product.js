"use strict";
(function () {
    $('.form').submit(function(e) {
        e.preventDefault(); // avoid to execute the actual submit of the form.
        $('.product-item').remove()
        let form = $(this);
        let url = form.attr('action');
        let productName=$("#searchProduct").val()
        let searchObj = {"product":productName}
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
                for (let value of result.products) {
                    $('#products').append('<div class="col-md-4 card mb-4 box-shadow product-item"><h4>'+ value.name +'</h4><p>Protein:'+value.nutrition.protein+'</p><p>Fat:'+value.nutrition.fat+'</p><p>Energy:'+value.nutrition.energy+'</p></div>')
                }
            },
            error: function(err) {
                errorSearch();
            },
        });
    });
}());

function errorSearch(){
    $('#errorSearch').show()
    setTimeout(() => $('#errorSearch').hide(), 3000);
}