let url = "http://comments.reillybrothers.net";

function put_comments(post_id, div) {
    const xmlhttp = new XMLHttpRequest();
    xmlhttp.withCredentials = false;
    xmlhttp.onload = function () {

        if (xmlhttp.status == 200) {
            const json = JSON.parse(xmlhttp.responseText).Comments;
            const h = document.createElement("h3");
            h.innerText = "Comments";
            div.appendChild(h);
            const label = document.createElement("label");
            label.setAttribute("for", "comment-input");
            label.innerText = "Add comment";
            div.appendChild(label);
            const i = document.createElement("input");
            i.setAttribute("id", "comment-input");
            i.setAttribute("type", "text");
            i.addEventListener("keyup", function(event) {
                if ((27 === event.which) || (13 === event.which)) {
                    event.preventDefault();
                    const comment = event.currentTarget.value;
                    const xmlhttp = new XMLHttpRequest();
                    xmlhttp.withCredentials = false;
                    xmlhttp.onload = function () {
                        if (xmlhttp.status == 200) {
                            while (div.hasChildNodes()) {
                                div.removeChild(div.lastChild);
                            }
                            put_comments(post_id, div);
                        }
                    };
                    xmlhttp.onerror = function () { alert("Failed to send comment");  }

                    xmlhttp.open('GET', url + '/add?post_id=' + post_id + "&comment=" + comment, true);
                    xmlhttp.send();

                    //this should delete value from the input
                    event.currentTarget.value = "";
                }
            });
            div.appendChild(i);
            for (const index in json) {
                const d = document.createElement("div");
                d.setAttribute("class", "comment-comment");
                const comment = json[index];
                d.innerText = comment;
                div.appendChild(d);
		if (index < json.length - 1) {
                    div.appendChild(document.createElement("hr"));
		}
            }
        }
        else if (xmlhttp.status == 400) {
            alert('There was an error 400');
        }
        else {
            alert('something else other than 200 was returned');
        }
    };
    xmlhttp.onerror = function () { alert("Failed to load comments");  }

    xmlhttp.open('GET', url + '/get?post_id=' + post_id, true);
    xmlhttp.send();
}

function auto_comment() {
    var x = document.getElementById("comment");
    if (x == null) {
        x = document.createElement("div");
        x.setAttribute("id", "comment");
        document.body.appendChild(x);
    }
    put_comments(window.location.href, x);
}
