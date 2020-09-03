class ScannMatcher:

    @classmethod
    async def create(cls):
        matcher = ScannMatcher()
        # TODO: download image data, calculate user interest signatures, create ScaNN
        return matcher

    def get_matches(self, user_id):
        # this is a stub
        return [user_id]
