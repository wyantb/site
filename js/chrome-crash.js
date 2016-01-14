(function() {
    function saveFile(name, fileContents) {
        var blob = new window.Blob([fileContents], {
            type: 'text/plain'
        });
        var url = window.URL.createObjectURL(blob);
        var link = window.document.createElement('a');
        var ev;

        link.href = url;
        link.download = name;
        link.textContent = 'save';
        link.click();

        // FF guard
        setTimeout(function() {
            window.URL.revokeObjectURL(url);
        }, 500);

        printMsg('Done!');
    }

    function requestAndSave(url) {
        var xhr = new window.XMLHttpRequest();
        xhr.onreadystatechange = function () {
            if (this.readyState === 4 && this.status === 200) {
                saveFile(url, new window.Uint8Array(this.response));
            }

        };
        xhr.open('GET', url);
        xhr.responseType = 'arraybuffer';
        xhr.send();
    }

    var btns = document.querySelectorAll('[data-dl]');
    Array.prototype.forEach.call(btns, function (btn) {
        btn.addEventListener('click', function () {
            var url = btn.getAttribute('data-dl');
            printMsg('Downloading...');
            requestAndSave(url);
        });
    });

    function printMsg(msg) {
        console.log(msg);
        var div = document.createElement('div');
        div.innerText = msg;
        document.querySelectorAll('body')[0].appendChild(div);
    }
}());
