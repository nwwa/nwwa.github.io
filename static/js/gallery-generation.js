        function load_member_list() {
            console.log("Loading member list");

            $.getJSON("/galleries_list.json", function(data) {
                var container = $("#gallery_list p");
                container.empty();

                var user_list = document.createElement("ul");
                for (user of data.users) {
                    var name = document.createTextNode(user.name);

                    var list_link = document.createElement("a");
                    list_link.appendChild(name);
                    list_link.href = "/gallery.html#" + user.folder;

                    var list_item = document.createElement("li");
                    list_item.appendChild(list_link);
                    user_list.appendChild(list_item);
                }

                container.append(user_list);

                $("#lightgallery").hide();
                $("#gallery_list").show();
            });
        }

        function load_member_images(member_name) {
            console.log("Loading contents for user " + member_name);

            var url = "/galleries/" + member_name + "/images.json";
            $.getJSON(url, function(data) {
                var lightgallery = $("#lightgallery");
                lightgallery.empty()

                for (image of data.images) {
                    var img = document.createElement("img");
                    img.src = "galleries/" + member_name + "/thumbs/" + image.filename;
                    img.className = "gallery_thumb";

                    var link = document.createElement("a");
                    link.href = "galleries/" + member_name + "/images/" + image.filename;
                    link.className = "gallery_link";
                    link.appendChild(img);

                    lightgallery.append(link);
                }

                lightGallery(document.getElementById('lightgallery'), {
                    download: false,
                    thumbnail: true,
                });
                $("#gallery_list").hide();
                $("#lightgallery").show();
            });
        }

        function load_content() {
            var member_name = window.location.hash.substr(1);
            if (member_name === "") {
                load_member_list();
            } else {
                load_member_images(member_name);
            }
        }

        window.onhashchange = load_content;
        load_content();
