# --- Stage 1: Build stage with full Python environment ---
FROM python:3.11-slim AS builder

# Set the working directory
WORKDIR /app

# Create and activate a virtual environment
RUN python -m venv /app/.venv
ENV PATH="/app/.venv/bin:$PATH"

# Copy and install dependencies
# This leverages layer caching. It only re-runs if requirements.txt changes.
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


# --- Stage 2: Final production stage ---
FROM python:3.11-slim

# Set the working directory
WORKDIR /app

# Copy the installed dependencies from the virtual environment in the builder stage
COPY --from=builder /app/.venv ./.venv

# Set the PATH to include the venv's bin directory
ENV PATH="/app/.venv/bin:$PATH"

# Copy the application source code
COPY ./app ./app

# Expose the port the app runs on
EXPOSE 8000

# Command to run the application using Uvicorn
# --host 0.0.0.0 makes the server accessible from outside the container
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]