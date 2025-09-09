components {
  id: "player"
  component: "/main/player/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "playback_rate: 0.1\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/player/newplayer.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 2.5
    y: 2.5
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"floor\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -18.0\n"
  "      y: -20.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"player\"\n"
  "  }\n"
  "  data: 26.812\n"
  "  data: 32.486427\n"
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
  id: "attack"
  type: "sound"
  data: "sound: \"/assets/sounds/slash.ogg\"\n"
  ""
}
embedded_components {
  id: "attack2"
  type: "sound"
  data: "sound: \"/assets/sounds/attack.ogg\"\n"
  ""
}
embedded_components {
  id: "hit"
  type: "sound"
  data: "sound: \"/assets/sounds/hit.ogg\"\n"
  ""
}
embedded_components {
  id: "duckcollision"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -7.0\n"
  "      y: -30.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.1234035\n"
  "      w: 0.9923566\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"duck\"\n"
  "  }\n"
  "  data: 26.30775\n"
  "  data: 23.388447\n"
  "  data: 1167.5427\n"
  "}\n"
  ""
}
