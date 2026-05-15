# Newton SDK feature factory

from feature.base_feature import NewtonBaseFeature
from feature.test_feature import NewtonTestFeature


def _make_feature(name):
    features = {
        "base": lambda: NewtonBaseFeature(),
        "test": lambda: NewtonTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
