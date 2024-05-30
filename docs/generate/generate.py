import os
import requests

def extract_paths(oas3_schema):
    """
    Extracts and groups paths by their associated API tags from an OAS3 schema.

    :param oas3_schema: The OAS3 schema as a dictionary.
    :return: A dictionary with API tags as keys and lists of paths as values.
    """
    paths = oas3_schema.get("paths", {})
    grouped_paths = {}

    for path, methods in paths.items():
        for method, details in methods.items():
            api_tag = details.get("tags", ["default"])[0]
            if api_tag not in grouped_paths:
                grouped_paths[api_tag] = []
            grouped_paths[api_tag].append({
                "path": path,
                "method": method.upper(),
                "summary": details.get("summary", ""),
                "description": details.get("description", ""),
            })
    
    return grouped_paths

def generate_markdown_table(grouped_paths):
    """
    Generates GitHub-flavored Markdown tables for each API tag.

    :param grouped_paths: A dictionary with API tags as keys and lists of paths as values.
    :return: A dictionary with API tags as keys and Markdown table strings as values.
    """
    markdown_tables = {}
    
    for api, paths in grouped_paths.items():
        markdown = []
        markdown.append(f"## {api} API\n")
        markdown.append("| Path | Method | Summary | Description | go-api-sdk-jamfpro coverage |")
        markdown.append("|------|--------|---------|-------------|-----------------------------|")
        for path_info in paths:
            markdown.append(
                f"| {path_info['path']} | {path_info['method']} | {path_info['summary']} | {path_info['description']} |  |"
            )
        markdown.append("\n")
        markdown_tables[api] = "\n".join(markdown)
    
    return markdown_tables

def save_to_markdown(markdown_tables, output_dir):
    """
    Saves the generated Markdown content to separate files for each API tag.

    :param markdown_tables: A dictionary with API tags as keys and Markdown table strings as values.
    :param output_dir: The directory to save the Markdown files to.
    """
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    for api, markdown in markdown_tables.items():
        filename = os.path.join(output_dir, f"{api.replace('/', '_')}.md")
        with open(filename, "w", encoding="utf-8") as file:
            file.write(markdown)
        print(f"Markdown file saved to {filename}")

def main():
    """
    Main function to load the OAS3 schema, extract paths, generate Markdown, and save it to files.
    """
    # Load the OAS3 schema
    schema_url = "https://lbgsandbox.jamfcloud.com/api/schema/"
    response = requests.get(schema_url, timeout=10)
    oas3_schema = response.json()

    # Extract and group paths
    grouped_paths = extract_paths(oas3_schema)

    # Generate Markdown tables
    markdown_tables = generate_markdown_table(grouped_paths)

    # Save to separate GitHub-based Markdown files for each API tag
    output_dir = "api_paths"
    save_to_markdown(markdown_tables, output_dir)

if __name__ == "__main__":
    main()
