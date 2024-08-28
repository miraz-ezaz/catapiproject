<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/static/css/home.css">
    <title>Voting</title>
</head>
<body>
    <div class="navbar">
        <a href="/" class="active">Voting</a>
        <a href="/breeds">Breeds</a>
        <a href="/favorites">Favs</a>
    </div>

    <div class="image-container">
        {{if .ImageURL}}
            <img src="{{.ImageURL}}" alt="Cat Image">
        {{else}}
            <p>No cat image available.</p>
        {{end}}
    </div>

    <div class="actions">
        <form method="post">
            <input type="hidden" name="image_id" value="{{.ImageID}}">
            <button type="submit" name="action" value="fav">❤️</button>
            <button type="submit" name="action" value="like">👍</button>
            <button type="submit" name="action" value="dislike">👎</button>
        </form>
    </div>
</body>
</html>
