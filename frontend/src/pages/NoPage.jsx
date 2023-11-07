import React from "react";


export default function NoPage() {
    return(
        <div class="container text-center mt-5">
            <div class="alert alert-danger">
                <h3 class="display-4">404 - Page Not Found</h3>
                <p class="lead">Sorry, the page you are looking for does not exist.</p>
                <a href="/" class="btn btn-primary">Go to Homepage</a>
            </div>
        </div>
    )
}