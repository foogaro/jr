package constants

import "time"

const NUM = 1
const LOCALE = "US"
const FREQUENCY = -1
const DURATION = 0
const DEFAULT_HOMEDIR = "$HOME/.jr"
const TEMPLATEDIR = "$HOME/.jr/templates"
const DEFAULT_KEY = "null"
const DEFAULT_OUTPUT = "stdout"
const DEFAULT_OUTPUT_TEMPLATE = "{{.V}}\n"
const DEFAULT_OUTPUT_KCAT_TEMPLATE = "{{.K}},{{.V}}\n"
const DEFAULT_SERIALIZER = "json-schema"
const KAFKA_CONFIG = "./kafka/config.properties"
const REGISTRY_CONFIG = "./kafka/registry.properties"
const REDIS_TTL = 1 * time.Minute
const REDIS_CONFIG = "./redis/config.json"
const MONGO_CONFIG = "./mongoDB/config.json"
const NUM_TEMPLATES = 5
const DEFAULT_PARTITIONS = 6
const DEFAULT_REPLICA = 3
const DEFAULT_PRELOAD_SIZE = 0
const DEFAULT_ENV_PREFIX = "JR"
const DEFAULT_EMITTER_NAME = "emitter"
const DEFAULT_VALUE_TEMPLATE = "user"
const DEFAULT_TOPIC = "test"