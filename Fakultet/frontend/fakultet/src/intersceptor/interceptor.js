// src/interceptor.js
const customFetch = async (url, options = {}) => {
    const token = localStorage.getItem('accessToken');
    console.log("Token: ", token);

    // Add common headers or other configurations here
    const modifiedOptions = {
        ...options,
        headers: {
            ...options.headers,
            'Authorization': `Bearer ${token}`, // Use the token from local storage
            'Content-Type': 'application/json', // Add default Content-Type header
        },
    };

    console.log("Intercepted request: ", modifiedOptions);

    try {
        const response = await fetch(url, modifiedOptions);

        // Handle common response scenarios
        if (!response.ok) {
            // Handle errors
            const errorData = await response.json();
            throw new Error(errorData.message || 'Something went wrong');
        }

        return response.json();
    } catch (error) {
        // Handle common errors
        console.error('Fetch error:', error.message);
        throw error;
    }
};

export default customFetch;
