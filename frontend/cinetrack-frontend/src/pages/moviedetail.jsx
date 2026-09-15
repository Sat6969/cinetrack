import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import { Getmoviebyid } from "../services/movieService";
import { senddata } from "../services/senddata";
import { getTracking } from "../services/trackingService";
import { getToken } from "../services/authService";

function Moviedetail() {
  const { id } = useParams();

  const [movie, setmovie] = useState(null);

  const [loading, setloading] = useState(true);
  const [error, seterror] = useState("");

  const [status, setstatus] = useState("");
  const [userRating, setUserRating] = useState("");
  const [review, setreview] = useState("");

  const [saving, setSaving] = useState(false);
  const [saveError, setSaveError] = useState("");
  const [saveSuccess, setSaveSuccess] = useState("");

  useEffect(() => {
    async function loadMovie() {
      try {
        // public movie information
        const movieData = await Getmoviebyid(id);

        setmovie(movieData);

        // tracking sirf logged-in user ke liye
        const token = getToken();

        if (token) {
          try {
            const trackingData = await getTracking(id);

            setstatus(trackingData.status || "");

            if (trackingData.userRating !== null) {
              setUserRating(trackingData.userRating);
            }

            setreview(trackingData.review || "");
          } catch (err) {
            console.log("could not load tracking");
          }
        }
      } catch (err) {
        seterror("something went wrong");
      } finally {
        setloading(false);
      }
    }

    loadMovie();
  }, [id]);

  async function handleSave() {
    const token = getToken();

    if (!token) {
      setSaveError("please login first");
      return;
    }

    if (status === "") {
      setSaveError("please select a status");
      return;
    }

    if (userRating === "") {
      setSaveError("please enter a rating");
      return;
    }

    const ratingNumber = Number(userRating);

    if (ratingNumber < 1 || ratingNumber > 10) {
      setSaveError("rating must be between 1 and 10");
      return;
    }

    const userreview = {
      status: status,
      userRating: ratingNumber,
      review: review,
    };

    try {
      setSaving(true);
      setSaveError("");
      setSaveSuccess("");

      await senddata(userreview, id);

      setSaveSuccess("saved successfully");
    } catch (err) {
      setSaveError(err.message);
    } finally {
      setSaving(false);
    }
  }

  if (loading === true) {
    return <div>loading...</div>;
  }

  if (error !== "") {
    return <div>{error}</div>;
  }

  if (!movie) {
    return null;
  }

  return (
    <div className="movie-detail">
      <img
        src={movie.poster}
        alt={movie.title}
        className="poster"
      />

      <div className="details">
        <div className="title">
          {movie.title}
        </div>

        <div className="year-rating">
          <div className="year">
            {movie.year}
          </div>

          <div className="rating">
            {movie.rating}
          </div>
        </div>

        <div className="genre">
          {movie.genres.map((genre) => (
            <span key={genre}>
              {genre}
            </span>
          ))}
        </div>
      </div>

      <div className="status-section">
        <select
          value={status}
          onChange={(e) => {
            setstatus(e.target.value);
          }}
        >
          <option value="">
            Select status
          </option>

          <option value="Watched">
            Watched
          </option>

          <option value="Watching">
            Watching
          </option>

          <option value="Want to Watch">
            Want to Watch
          </option>

          <option value="Dropped">
            Dropped
          </option>
        </select>
      </div>

      <div className="rating-section">
        <label>Your rating</label>

        <input
          type="number"
          min="1"
          max="10"
          value={userRating}
          onChange={(e) => {
            setUserRating(e.target.value);
          }}
        />
      </div>

      <div className="review-section">
        <label>Your review</label>

        <textarea
          value={review}
          placeholder="Write your review..."
          onChange={(e) => {
            setreview(e.target.value);
          }}
        />
      </div>

      <button
        onClick={handleSave}
        disabled={saving}
      >
        {saving ? "Saving..." : "Save"}
      </button>

      {saveError !== "" && (
        <div>{saveError}</div>
      )}

      {saveSuccess !== "" && (
        <div>{saveSuccess}</div>
      )}
    </div>
  );
}

export default Moviedetail;