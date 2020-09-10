import unittest
import asyncio

import numpy as np

from .scann_matcher import ScannMatcher

class MockImageLoader:

    def __init__(self, images=[]):
        self.images = images

    def get_images(self):
        return self.images


class TestScannMatcher(unittest.TestCase):

    def test_returns_empty_list_for_user_with_no_photos(self):

        image_loader = MockImageLoader([
            { "owner_id": "user1", "labels": [{ "name": "label1", "score": 0.9 }] },
            { "owner_id": "user2", "labels": [{ "name": "label2", "score": 0.8 }] },
        ])

        matcher = asyncio.run(ScannMatcher.create(image_loader=image_loader))

        actual = matcher.get_matches("user3")

        self.assertEqual(actual, [])

    def test_filters_labels_with_low_score(self):

        image_loader = MockImageLoader([
            {
                "owner_id": "user1",
                "labels": [
                    { "name": "label1", "score": 0.95 },
                    { "name": "label2", "score": 0.85 },
                    { "name": "label3", "score": 0.75 },
                ]
            },
        ])

        matcher = asyncio.run(ScannMatcher.create(image_loader=image_loader, min_score=0.8))

        actual = matcher.interest_signatures.shape[1] # amount of selected labels

        self.assertEqual(actual, 2)

    def test_selects_most_popular_labels(self):

        image_loader = MockImageLoader([
            {
                "owner_id": "user1",
                "labels": [
                    { "name": "label1", "score": 0.95 },
                    { "name": "label2", "score": 0.85 },
                    { "name": "label3", "score": 0.75 },
                ]
            },
            {
                "owner_id": "user2",
                "labels": [
                    { "name": "label1", "score": 0.95 },
                    { "name": "label2", "score": 0.85 },
                    { "name": "label4", "score": 0.75 },
                ]
            },
            {
                "owner_id": "user3",
                "labels": [
                    { "name": "label1", "score": 0.95 },
                    { "name": "label5", "score": 0.75 },
                ]
            },
        ])

        matcher = asyncio.run(ScannMatcher.create(image_loader=image_loader, max_features=2))

        user3_signature = matcher.interest_signatures[matcher.user_idx["user3"]]

        self.assertEqual(list(user3_signature), [1, 0])

if __name__ == '__main__':
    unittest.main()
