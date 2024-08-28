document.getElementById("breed-dropdown").addEventListener("change", function() {
    var breedID = this.value;
    var xhr = new XMLHttpRequest();

    xhr.open("POST", "/breeds", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var response = JSON.parse(xhr.responseText);

            // Update Breed Details
            var breedDetails = response.breedDetails;
            if (breedDetails) {
                var detailsHTML = `
                    <h2>${breedDetails.name}</h2>
                    <p>${breedDetails.origin} (${breedDetails.id})</p>
                    <p>${breedDetails.description}</p>
                    <a href="${breedDetails.wikipedia_url}" target="_blank">Wikipedia</a>
                `;
                document.getElementById("breed-details").innerHTML = detailsHTML;
            }

            // Update Breed Images
            var breedImages = response.breedImages;
            if (breedImages) {
                var carouselHTML = breedImages.map(image => `<img src="${image.url}" alt="Breed Image">`).join('');
                document.getElementById("breed-carousel").innerHTML = carouselHTML;

                // Reset and start the slideshow after the new images are loaded
                resetSlideshow();
                startSlideshow();
            }
        }
    };
    xhr.send("breed_id=" + breedID);
});

let slideshowInterval;

function resetSlideshow() {
    const carouselSlide = document.querySelector('.carousel-slide');
    carouselSlide.style.transform = 'translateX(0px)'; // Reset the transform to the initial state
    clearInterval(slideshowInterval); // Clear any existing intervals to prevent overlapping animations
}

function startSlideshow() {
    const carouselSlide = document.querySelector('.carousel-slide');
    const images = document.querySelectorAll('.carousel-slide img');
    let counter = 0;
    const size = images[0].clientWidth;

    slideshowInterval = setInterval(() => {
        counter++;
        if (counter >= images.length) {
            counter = 0;
        }
        carouselSlide.style.transition = 'transform 0.5s ease-in-out';
        carouselSlide.style.transform = 'translateX(' + (-size * counter) + 'px)';
    }, 3000); // Change image every 3 seconds
}

// Initialize the slideshow on page load
document.addEventListener("DOMContentLoaded", function() {
    startSlideshow();
});