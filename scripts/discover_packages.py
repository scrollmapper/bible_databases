import os
import ast
import sys

def get_imported_packages(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        node = ast.parse(file.read(), file_path)
    return {alias.name for n in ast.walk(node) if isinstance(n, ast.Import) for alias in n.names}

def get_imported_from_packages(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        node = ast.parse(file.read(), file_path)
    return {n.module for n in ast.walk(node) if isinstance(n, ast.ImportFrom)}

def is_standard_package(package_name):
    standard_libs = set(sys.stdlib_module_names)
    return package_name.split('.')[0] in standard_libs

def main(directory):
    all_packages = set()
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith('.py'):
                file_path = os.path.join(root, file)
                imported_packages = get_imported_packages(file_path)
                imported_from_packages = get_imported_from_packages(file_path)
                all_packages.update(imported_packages)
                all_packages.update(imported_from_packages)
    
    # Filter out standard packages
    non_standard_packages = {pkg for pkg in all_packages if pkg and not is_standard_package(pkg)}
    
    if non_standard_packages:
        print("You need to install the following packages:")
        for package in sorted(non_standard_packages):
            print(f"pip install {package}")
    else:
        print("No non-standard packages found.")

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python discover_packages.py <directory>")
        sys.exit(1)
    
    directory = sys.argv[1]
    main(directory)
