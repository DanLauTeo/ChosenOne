from google.cloud import datastore

class ImageLoader:
    def get_images(self):
        datastore_client = datastore.Client()

        query = datastore_client.query(kind="Image")
        query.add_filter("type", "=", "user_uploaded_image")

        return list(query.fetch())
