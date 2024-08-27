<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Image</title>
</head>
<body>
    <h1>Random Cat Image</h1>
    {{if .ImageURL}}
        <img src="{{.ImageURL}}" alt="Cat Image" style="max-width:100%; height:auto;">
    {{else}}
        <p>No image available.</p>
    {{end}}
</body>
</html>
