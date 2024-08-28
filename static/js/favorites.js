const gallery = document.getElementById('gallery');
const gridViewBtn = document.getElementById('grid-view');
const listViewBtn = document.getElementById('list-view');

gridViewBtn.addEventListener('click', function() {
    gallery.classList.remove('list-view');
    gallery.classList.add('gallery');
    gridViewBtn.classList.add('active');
    listViewBtn.classList.remove('active');
});

listViewBtn.addEventListener('click', function() {
    gallery.classList.add('list-view');
    gridViewBtn.classList.remove('active');
    listViewBtn.classList.add('active');
});