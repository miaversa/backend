(function ($) {
    $.fn.pagseguro = function() {

        function getSessionID() {
            $.get("/pagseguro/session")
            .done(function(data) {
                PagSeguroDirectPayment.setSessionId(data);
            })
            .fail(function() {
                alert("error");
            });
        }

        getSessionID();
    };
}(jQuery));
