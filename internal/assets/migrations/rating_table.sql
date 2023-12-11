CREATE TABLE rating (
    id SERIAL PRIMARY KEY,
    review_id INT NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    FOREIGN KEY (review_id) REFERENCES reviews(id)
);