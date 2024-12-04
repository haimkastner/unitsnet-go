from common.fetch_units_definitions import get_definitions
from generators.generate_unit_class import unit_class_generator
from generators.generate_docs import units_docs_generator

print("Starting generating python units...")

# Fetch all units definitions
definitions = get_definitions(repo_owner_and_name="angularsen/UnitsNet")

# Generate python unit class for each unit definition
for definition in definitions:
    unit_class_generator(unit_definition=definition)

# # Generate README doc file
units_docs_generator(definitions)

print("Generating python units package finished successfully")
