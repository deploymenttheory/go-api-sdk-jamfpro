#!/bin/bash
# Script to generate a .p12 certificate for testing purposes on macOS

# Set default values
CERT_NAME="test_certificate"
OUTPUT_DIR="$(pwd)"
DAYS_VALID=365
KEY_SIZE=2048
COUNTRY="US"
STATE="California"
LOCALITY="San Francisco"
ORGANIZATION="Test Organization"
ORG_UNIT="Testing"
CN="localhost"

# Display help information
show_help() {
    echo "Usage: $0 [options]"
    echo "Generate a .p12 certificate for testing purposes."
    echo ""
    echo "Options:"
    echo "  -n, --name NAME       Certificate name (default: $CERT_NAME)"
    echo "  -o, --output DIR      Output directory (default: current directory)"
    echo "  -d, --days DAYS       Days the certificate is valid (default: $DAYS_VALID)"
    echo "  -k, --keysize SIZE    Key size in bits (default: $KEY_SIZE)"
    echo "  -c, --country CODE    Country code (default: $COUNTRY)"
    echo "  -s, --state STATE     State or province (default: $STATE)"
    echo "  -l, --locality LOC    Locality/city (default: $LOCALITY)"
    echo "  -g, --org ORG         Organization (default: $ORGANIZATION)"
    echo "  -u, --unit UNIT       Organizational unit (default: $ORG_UNIT)"
    echo "  --cn CN               Common Name (default: $CN)"
    echo "  -h, --help            Display this help message"
    exit 0
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        -n|--name)
            CERT_NAME="$2"
            shift 2
            ;;
        -o|--output)
            OUTPUT_DIR="$2"
            shift 2
            ;;
        -d|--days)
            DAYS_VALID="$2"
            shift 2
            ;;
        -k|--keysize)
            KEY_SIZE="$2"
            shift 2
            ;;
        -c|--country)
            COUNTRY="$2"
            shift 2
            ;;
        -s|--state)
            STATE="$2"
            shift 2
            ;;
        -l|--locality)
            LOCALITY="$2"
            shift 2
            ;;
        -g|--org)
            ORGANIZATION="$2"
            shift 2
            ;;
        -u|--unit)
            ORG_UNIT="$2"
            shift 2
            ;;
        --cn)
            CN="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use --help for usage information."
            exit 1
            ;;
    esac
done

# Ensure output directory exists
mkdir -p "$OUTPUT_DIR"

# Set file paths
KEY_FILE="$OUTPUT_DIR/${CERT_NAME}.key"
CSR_FILE="$OUTPUT_DIR/${CERT_NAME}.csr"
CERT_FILE="$OUTPUT_DIR/${CERT_NAME}.crt"
DER_FILE="$OUTPUT_DIR/${CERT_NAME}.der"
P12_FILE="$OUTPUT_DIR/${CERT_NAME}.p12"
P12_FILE="$OUTPUT_DIR/${CERT_NAME}.p12"

echo "Generating private key..."
openssl genrsa -out "$KEY_FILE" $KEY_SIZE

echo "Creating certificate signing request..."
openssl req -new -key "$KEY_FILE" -out "$CSR_FILE" -subj "/C=$COUNTRY/ST=$STATE/L=$LOCALITY/O=$ORGANIZATION/OU=$ORG_UNIT/CN=$CN"

echo "Creating self-signed certificate..."
openssl x509 -req -days $DAYS_VALID -in "$CSR_FILE" -signkey "$KEY_FILE" -out "$CERT_FILE"

echo "Converting certificate to DER format..."
openssl x509 -in "$CERT_FILE" -outform DER -out "$DER_FILE"

echo "Creating PKCS#12 (.p12) file for macOS Keychain import..."
echo "You will be prompted to set a password for the certificate:"
openssl pkcs12 -export -out "$P12_FILE" -inkey "$KEY_FILE" -in "$CERT_FILE"

echo "Certificate generation complete!"
echo "Files created:"
echo "  Private key: $KEY_FILE"
echo "  Certificate signing request: $CSR_FILE"
echo "  Certificate (PEM): $CERT_FILE"
echo "  Certificate (DER): $DER_FILE"
echo "  PKCS#12 bundle: $P12_FILE"

echo ""
echo "To import the certificate into macOS Keychain:"
echo "  open $P12_FILE"
echo "  (You'll need to enter the password you just created)"
echo ""
echo "To use the .p12 certificate for testing:"
echo "  The certificate is available at: $P12_FILE"
echo "  (Remember to use the password you set when prompted by applications)"