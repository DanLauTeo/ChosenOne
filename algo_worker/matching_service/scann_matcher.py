import numpy as np
import scipy.spatial.distance as distance
import heapq

class ScannMatcher:

    @classmethod
    async def create(cls, image_loader, min_score=0.7, max_features=256):
        matcher = ScannMatcher()

        images = image_loader.get_images()

        labels = matcher.get_features(images, min_score=min_score, max_features=max_features)

        users = matcher.get_active_users(images)

        user_idx = index_dict(users)
        label_idx = index_dict(labels)

        matcher.users = users
        matcher.user_idx = user_idx
        matcher.interest_signatures = matcher.get_interest_signatures(images, user_idx, label_idx)

        return matcher

    def get_features(self, images, min_score, max_features):
        label_frequency = {}

        for image in images:
            for label in filter(lambda x: x["score"] > min_score, image["labels"]):
                name = label["name"]
                label_frequency[name] = label_frequency.get(name, 0) + 1

        labels = heapq.nlargest(
            min(len(label_frequency), max_features),
            label_frequency.items(),
            lambda x: x[1]
        )

        return list(map(lambda x:x[0], labels))

    def get_active_users(self, images):
        return list(set([image["owner_id"] for image in images]))

    def get_interest_signatures(self, images, user_idx, label_idx):
        interest_signatures = np.zeros(shape=(len(user_idx), len(label_idx)))

        for image in images:

            user_signature = interest_signatures[user_idx[image["owner_id"]]]

            for label in filter(lambda x: x["name"] in label_idx, image["labels"]):
                user_signature[label_idx[label["name"]]] += 1

        interest_signatures /= np.linalg.norm(interest_signatures, axis=1)[:, np.newaxis]

        return interest_signatures


    def get_matches(self, user_id, max_results = 100):
        try:
            user_index = self.user_idx[user_id]
            user_signature = self.interest_signatures[user_index]
        except KeyError:
            return []

        # uses bruteforce implementation temporarily

        heap = []

        for idx, signature in enumerate(self.interest_signatures):
            if idx == user_index:
                continue

            item = distance.cosine(signature, user_signature), idx

            if len(heap) < max_results:
                heapq.heappush(heap, item)
            else:
                heapq.heappushpop(heap, item)

        return list(map(lambda x: self.users[x[1]], heap))


def index_dict(lst):
    index_of = dict.fromkeys(lst)

    for i, value in enumerate(lst):
        index_of[value] = i

    return index_of
