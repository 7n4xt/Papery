document.addEventListener('DOMContentLoaded', function() {
    const searchForm = document.querySelector('.search-form');
    const orientationFilter = document.getElementById('orientation-filter');
    const sizeFilter = document.getElementById('size-filter');
    const colorFilter = document.getElementById('color-filter');

    // Add filters to form submission
    searchForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const searchQuery = this.querySelector('input[name="search-word"]').value;
        let url = `/search?search-word=${encodeURIComponent(searchQuery)}`;

        // Add selected filters to URL
        if (orientationFilter.value) {
            url += `&orientation=${orientationFilter.value}`;
        }
        if (sizeFilter.value) {
            url += `&size=${sizeFilter.value}`;
        }
        if (colorFilter.value) {
            url += `&color=${colorFilter.value}`;
        }

        window.location.href = url;
    });

    // Update filters when changed
    [orientationFilter, sizeFilter, colorFilter].forEach(filter => {
        filter.addEventListener('change', function() {
            searchForm.dispatchEvent(new Event('submit'));
        });
    });
});