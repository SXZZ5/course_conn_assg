import { useEffect, useState } from "react";
export default function Results() {
    const [cookies, setCookies] = useState([]);

    useEffect(() => {
        const fetchCookies = async () => {
            try {
                const response = await fetch('http://localhost:8080/oldcookies', {
                    method: 'GET',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    }
                });

                if (response.ok) {
                    const data = await response.json();
                    setCookies(data.cookies);
                } else {
                    console.error('Failed to fetch cookies');
                }
            } catch (error) {
                console.error('Error:', error);
            }
        };

        fetchCookies();
    }, []);

    return (
        <div>
            <h2>Old Cookies</h2>
            {cookies.length > 0 ? (
                <ul>
                    {cookies.map((cookie, index) => (
                        <li key={index}>{cookie}</li>
                    ))}
                </ul>
            ) : (
                <p>No cookies found.</p>
            )}
        </div>
    );
}