import os
import re

# Define the base folder to scan
base_domain_folder = r"c:\Users\Jimmy Castillo\earnforglance\src\server\api\controller"
output_file = "class_diagram.puml"

# Regex to match struct declarations
struct_regex = re.compile(r"struct {")

# Dictionary to store structs and their relationships
structs = {}
relationships = []

def extract_structs_and_relationships(file_path):
    with open(file_path, "r", encoding="utf-8") as file:
        content = file.read()
        matches = struct_regex.findall(content)
       
        print(f"{file_path}")
        for match in matches:
            structs[match] = []  # Initialize with no relationships


        # Find relationships (e.g., struct fields referencing other structs)
        for struct_name in structs.keys():
            if struct_name in content:
                relationships.append((os.path.basename(file_path), struct_name))

        

# Walk through the base domain folder and its subfolders
for root, _, files in os.walk(base_domain_folder):
   
    for file in files:       
        if file.endswith(".go"):  # Assuming Go files            
            file_path = os.path.join(root, file)            
            extract_structs_and_relationships(file_path)

# Generate PlantUML content
with open(output_file, "w", encoding="utf-8") as puml_file:
    puml_file.write("@startuml\n\n")
    # Write class definitions
    for struct_name in structs.keys():
        puml_file.write(f"class {struct_name} {{\n}}\n")
    # Write relationships
    for file_name, struct_name in relationships:
        puml_file.write(f"{file_name} --> {struct_name}\n")
    puml_file.write("\n@enduml")

print(f"PlantUML file generated: {output_file}")