CREATE TABLE application (
    uuid TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    life_id INT NOT NULL,
    charm_uuid TEXT NOT NULL,
    charm_modified_version INT,
    charm_upgrade_on_error BOOLEAN DEFAULT FALSE,
    exposed BOOLEAN DEFAULT FALSE,
    placement TEXT,
    password_hash_algorithm_id TEXT,
    password_hash TEXT,
    CONSTRAINT fk_application_life
    FOREIGN KEY (life_id)
    REFERENCES life (id),
    CONSTRAINT fk_application_charm
    FOREIGN KEY (charm_uuid)
    REFERENCES charm (uuid),
    CONSTRAINT fk_application_password_hash_algorithm
    FOREIGN KEY (password_hash_algorithm_id)
    REFERENCES password_hash_algorithm (id)
);

CREATE UNIQUE INDEX idx_application_name
ON application (name);

CREATE TABLE cloud_service (
    uuid TEXT NOT NULL PRIMARY KEY,
    application_uuid TEXT NOT NULL,
    net_node_uuid TEXT NOT NULL,
    provider_id TEXT NOT NULL,
    CONSTRAINT fk_cloud_service_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_cloud_service_net_node
    FOREIGN KEY (net_node_uuid)
    REFERENCES net_node (uuid)
);

CREATE UNIQUE INDEX idx_cloud_service_provider
ON cloud_service (provider_id);

CREATE INDEX idx_cloud_service_application
ON cloud_service (application_uuid);

CREATE UNIQUE INDEX idx_cloud_service_net_node
ON cloud_service (net_node_uuid);

-- Application scale is currently only targeting k8s applications.
CREATE TABLE application_scale (
    application_uuid TEXT NOT NULL PRIMARY KEY,
    scale INT,
    scale_target INT,
    scaling BOOLEAN DEFAULT FALSE,
    CONSTRAINT fk_application_endpoint_scale_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid)
);

CREATE TABLE application_endpoint_space (
    application_uuid TEXT NOT NULL,
    space_uuid TEXT,
    CONSTRAINT fk_application_endpoint_space_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_application_endpoint_space_space
    FOREIGN KEY (space_uuid)
    REFERENCES space (uuid),
    PRIMARY KEY (application_uuid, space_uuid)
);

CREATE TABLE application_endpoint_cidr (
    application_uuid TEXT NOT NULL,
    cidr TEXT,
    CONSTRAINT fk_application_endpoint_cidr_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    PRIMARY KEY (application_uuid, cidr)
);

CREATE TABLE application_config_hash (
    application_uuid TEXT NOT NULL PRIMARY KEY,
    sha256 TEXT NOT NULL,
    CONSTRAINT fk_application_config_hash_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid)
);

CREATE TABLE application_config (
    application_uuid TEXT NOT NULL,
    "key" TEXT NOT NULL,
    type_id INT NOT NULL,
    value TEXT,
    CONSTRAINT fk_application_config_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_application_config_charm_config_type
    FOREIGN KEY (type_id)
    REFERENCES charm_config_type (id),
    PRIMARY KEY (application_uuid, "key")
);

CREATE VIEW v_application_config AS
SELECT
    a.uuid,
    ac."key",
    ac.value,
    cct.name AS type
FROM application AS a
LEFT JOIN application_config AS ac ON a.uuid = ac.application_uuid
INNER JOIN charm_config_type AS cct ON ac.type_id = cct.id;

CREATE TABLE application_constraint (
    application_uuid TEXT NOT NULL PRIMARY KEY,
    constraint_uuid TEXT,
    CONSTRAINT fk_application_constraint_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_application_constraint_constraint
    FOREIGN KEY (constraint_uuid)
    REFERENCES "constraint" (uuid)
);

CREATE TABLE application_setting (
    application_uuid TEXT NOT NULL PRIMARY KEY,
    trust BOOLEAN DEFAULT FALSE,
    CONSTRAINT fk_application_setting_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid)
);

CREATE TABLE application_platform (
    application_uuid TEXT NOT NULL,
    os_id TEXT NOT NULL,
    channel TEXT,
    architecture_id INT NOT NULL,
    CONSTRAINT fk_application_platform_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_application_platform_os
    FOREIGN KEY (os_id)
    REFERENCES os (id),
    CONSTRAINT fk_application_platform_architecture
    FOREIGN KEY (architecture_id)
    REFERENCES architecture (id)
);

CREATE TABLE application_channel (
    application_uuid TEXT NOT NULL,
    track TEXT,
    risk TEXT,
    branch TEXT,
    CONSTRAINT fk_application_origin_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    PRIMARY KEY (application_uuid, track, risk, branch)
);

CREATE VIEW v_application_constraint AS
SELECT
    ac.application_uuid,
    c.arch,
    c.cpu_cores,
    c.cpu_power,
    c.mem,
    c.root_disk,
    c.root_disk_source,
    c.instance_role,
    c.instance_type,
    ctype.value AS container_type,
    c.virt_type,
    c.allocate_public_ip,
    c.image_id,
    ctag.tag,
    cspace.space,
    czone.zone
FROM application_constraint AS ac
INNER JOIN "constraint" AS c ON ac.constraint_uuid = c.uuid
LEFT JOIN container_type AS ctype ON c.container_type_id = ctype.id
LEFT JOIN constraint_tag AS ctag ON c.uuid = ctag.constraint_uuid
LEFT JOIN constraint_space AS cspace ON c.uuid = cspace.constraint_uuid
LEFT JOIN constraint_zone AS czone ON c.uuid = czone.constraint_uuid;
