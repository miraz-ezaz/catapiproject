<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Favorites</title>
    <link rel="stylesheet" href="/static/css/favorites.css">
</head>
<body>
    <div class="navbar">
        <a href="/">Voting</a>
        <a href="/breeds">Breeds</a>
        <a href="/favorites" class="active">Favs</a>
    </div>

    <div class="view-icons">
        <button id="grid-view" class="active">🔳</button>
        <button id="list-view">📋</button>
    </div>

    <div class="scrollable-container">
        <div id="gallery" class="gallery">
            {{range .FavoriteImages}}
            <img src="{{.URL}}" alt="Favorite Image">
            {{end}}
        </div>
    </div>
<script src="/static/js/favorites.js"></script>
</body>
</html>
