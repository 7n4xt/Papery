// Add this to your frontend JavaScript code
async function toggleFavorite(checkbox) {
    const isChecked = checkbox.checked;

    try {
        if (isChecked) {
            const photo = {
                id: parseInt(checkbox.dataset.photo),
                photographer: checkbox.dataset.photographer,
                alt: checkbox.dataset.alt,
                src: {
                    original: checkbox.dataset.srcOriginal,
                    large: checkbox.dataset.srcLarge,
                    large2x: checkbox.dataset.srcOriginal // Fallback to original
                }
            };
            
            await addToFavorites(photo);
        } else {
            await removeFromFavorites(parseInt(checkbox.dataset.photo));
        }
    } catch (error) {
        console.error('Error toggling favorite:', error);
        // Revert checkbox state on error
        checkbox.checked = !isChecked;
    }
}

async function addToFavorites(photo) {
    try {
        console.log('Adding to favorites:', photo); // Debug log
        const response = await fetch('/api/favorites/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(photo)
        });
        
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Failed to add to favorites: ${errorText}`);
        }
    } catch (error) {
        console.error('Error adding to favorites:', error);
        throw error;
    }
}

async function removeFromFavorites(photoId) {
    try {
        const response = await fetch(`/api/favorites/remove?id=${photoId}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Failed to remove from favorites: ${errorText}`);
        }
    } catch (error) {
        console.error('Error removing from favorites:', error);
        throw error;
    }
}

async function checkFavoriteStatus(photoId) {
    try {
        const response = await fetch('/favorite');
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Failed to fetch favorites: ${errorText}`);
        }
        
        const text = await response.text();
        const tempDiv = document.createElement('div');
        tempDiv.innerHTML = text;
        
        const favoritePhotos = Array.from(tempDiv.querySelectorAll('.photo-container'))
            .map(el => {
                const onclick = el.querySelector('.remove-btn').getAttribute('onclick');
                const match = onclick.match(/removeFavorite\('(\d+)'/);
                return match ? parseInt(match[1]) : null;
            })
            .filter(id => id !== null);
        
        document.querySelector('.like').checked = favoritePhotos.includes(parseInt(photoId));
    } catch (error) {
        console.error('Error checking favorite status:', error);
    }
}

// Function to handle removing favorites from the favorites page
function removeFavorite(photoId, button) {
    removeFromFavorites(photoId)
        .then(() => {
            // Remove the photo container from the grid
            const container = button.closest('.photo-container');
            if (container) {
                container.remove();
                
                // Check if there are any photos left
                const photosGrid = document.querySelector('.Photos-grid');
                if (photosGrid && photosGrid.children.length === 0) {
                    const photosContainer = document.querySelector('.photos');
                    photosContainer.innerHTML = `
                        <div class="empty-favorites">
                            <p>No favorite photos yet. Start adding some from the home page!</p>
                        </div>
                    `;
                }
            }
        })
        .catch(error => {
            console.error('Error removing favorite:', error);
        });
}