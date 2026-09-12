import { Getmoviebyid } from "../services/movieService";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { senddata } from "../services/senddata";

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
    async function getmovie() {
      try {
        const movie = await Getmoviebyid(id);
        setmovie(movie);
      } catch {
        seterror("something happend");
      } finally {
        setloading(false);
      }
    }
    getmovie();
  }, [id]);

  if (loading === true) {
    return <div>loading...</div>;
  }
  if (error !== "") {
    return <div>something went wrong</div>;
  }
  if (!movie) {
    return null;
  }
  async function handleSave() {
    const userreview = {
      status: status,
      userRating: Number(userRating),
      review: review,
    };

    try {
      setSaving(true);
      setSaveError("");
      setSaveSuccess("");

      const savedData = await senddata(userreview, id);

      setSaveSuccess("saved successfully");
      console.log(savedData);
    } catch {
      setSaveError("could not save");
    } finally {
      setSaving(false);
    }
  }
  return (
    <div className="movie-detail">
      <img src={movie.poster} alt={movie.title} className="poster" />
      <div className="details">
        <div className="title">{movie.title}</div>
        <div className="year-rating">
          <div className="year">{movie.year}</div>
          <div className="rating">{movie.rating}</div>
        </div>
        <div className="genre">
          {movie.genres.map((genre) => {
            return <span key={genre}>{genre}</span>;
          })}
        </div>
      </div>

      <div className="status-section">
        <select
          name=""
          id=""
          value={status}
          onChange={(e) => {
            setstatus(e.target.value);
          }}
        >
          <option value="">Select status</option>
          <option value="Watched">Watched</option>
          <option value="Watching">Watching</option>
          <option value="Want to Watch">Want to Watch</option>
          <option value="Dropped">Dropped</option>
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
      <button onClick={handleSave} disabled={saving}>
        {saving ? "Saving..." : "Save"}
      </button>
      {saveError !== "" && <div>{saveError}</div>}
      {saveSuccess !== "" && <div>{saveSuccess}</div>}
    </div>
  );
}

export default Moviedetail;
