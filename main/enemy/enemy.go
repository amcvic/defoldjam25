components {
  id: "player"
  component: "/main/enemy/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/player/player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
    z: 0.5
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: -26.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"enemy\"\n"
  "  }\n"
  "  data: 26.09134\n"
  "  data: 38.13117\n"
  "  data: 52.312\n"
  "}\n"
  ""
}
embedded_components {
  id: "jump"
  type: "sound"
  data: "sound: \"/assets/sounds/jump.ogg\"\n"
  ""
}
embedded_components {
  id: "hit"
  type: "sound"
  data: "sound: \"/assets/sounds/hit.ogg\"\n"
  ""
}
