components {
  id: "customer"
  component: "/main/customer/customer.script"
}
components {
  id: "progress_bar"
  component: "/main/customer/progress_bar.gui"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idleright\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/customer/customer.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "speech_bubble"
  type: "sprite"
  data: "default_animation: \"speechbubble\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/customer/food.atlas\"\n"
  "}\n"
  ""
  position {
    y: 109.0
  }
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "thinking_dots"
  type: "sprite"
  data: "default_animation: \"thinking\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/customer/food.atlas\"\n"
  "}\n"
  ""
  position {
    y: 175.0
    z: 1.0
  }
}
embedded_components {
  id: "food_factory"
  type: "factory"
  data: "prototype: \"/main/customer/food.go\"\n"
  ""
}
