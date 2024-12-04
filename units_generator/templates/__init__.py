

# Load jinja template for the Golang unit class
unit_class_template = ""
with open("units_generator/templates/unit_template.jinja2", "r", encoding="utf-8") as f:
    unit_class_template = f.read()


docs_template = ""
with open("units_generator/templates/docs_template.jinja2", "r", encoding="utf-8") as f:
    docs_template = f.read()

