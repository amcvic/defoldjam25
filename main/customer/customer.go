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
