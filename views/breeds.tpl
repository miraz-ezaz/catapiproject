<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Breeds</title>
    <link rel="stylesheet" href="/static/css/breeds.css">
</head>
<body>
    <div class="navbar">
        <a href="/" class="active">Voting</a>
        <a href="/breeds">Breeds</a>
        <a href="/favorites">Favs</a>
    </div>

    <div class="breed-selector">
        <select id="breed-dropdown">
            {{range .Breeds}}
            <option value="{{.ID}}">{{.Name}}</option>
            {{end}}
        </select>
    </div>

    <div class="carousel">
        <div id="breed-carousel" class="carousel-slide">
            {{if .BreedImages}}
            {{range .BreedImages}}
            <img src="{{.URL}}" alt="Breed Image">
            {{end}}
            {{else}}
            <p>No images available for this breed.</p>
            {{end}}
        </div>
    </div>

    <div class="breed-details" id="breed-details">
        {{if .BreedDetails}}
        <h2>{{.BreedDetails.Name}}</h2>
        <p>{{.BreedDetails.Origin}} ({{.BreedDetails.ID}})</p>
        <p>{{.BreedDetails.Description}}</p>
        <a href="{{.BreedDetails.WikipediaURL}}" target="_blank">Wikipedia</a>
        {{else}}
        <p>No details available for this breed.</p>
        {{end}}
    </div>

    <script src="/static/js/breeds.js"></script>
</body>
</html>
