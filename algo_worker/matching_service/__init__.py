from .image_loader import ImageLoader
from .scann_matcher import ScannMatcher


class MatchingService:

    @classmethod
    async def create(cls):
        service = MatchingService()
        await service.recalc_scann()
        return service

    def get_matches(self, user_id):
        return self.scann_matcher.get_matches(user_id)

    async def recalc_scann(self):

        image_loader = ImageLoader()

        scann_matcher = await ScannMatcher.create(image_loader=image_loader)

        self.scann_matcher = scann_matcher
