package storm

// #include <StormLib.h>
import "C"

const MAX_PATH uint32 = C.MAX_PATH

// Flags for SFileOpenArchive
const STREAM_PROVIDER_FLAT uint32 = C.STREAM_PROVIDER_FLAT       // Stream is linear with no offset mapping
const STREAM_PROVIDER_PARTIAL uint32 = C.STREAM_PROVIDER_PARTIAL // Stream is partial file (.part)
const STREAM_PROVIDER_MPQE uint32 = C.STREAM_PROVIDER_MPQE       // Stream is an encrypted MPQ
const STREAM_PROVIDER_BLOCK4 uint32 = C.STREAM_PROVIDER_BLOCK4   // 0x4000 per block, text MD5 after each block, max 0x2000 blocks per file

const BASE_PROVIDER_FILE uint32 = C.BASE_PROVIDER_FILE // Base data source is a file
const BASE_PROVIDER_MAP uint32 = C.BASE_PROVIDER_MAP   // Base data source is memory-mapped file
const BASE_PROVIDER_HTTP uint32 = C.BASE_PROVIDER_HTTP // Base data source is a file on web server

const STREAM_FLAG_READ_ONLY uint32 = C.STREAM_FLAG_READ_ONLY         // Stream is read only
const STREAM_FLAG_WRITE_SHARE uint32 = C.STREAM_FLAG_WRITE_SHARE     // This flag causes the writable MPQ being open for write share. Use with caution. If two applications write to an open MPQ simultaneously, the MPQ data get corrupted.
const STREAM_FLAG_USE_BITMAP uint32 = C.STREAM_FLAG_USE_BITMAP       // If the file has a file bitmap, load it and use it
const MPQ_OPEN_NO_LISTFILE uint32 = C.MPQ_OPEN_NO_LISTFILE           // Don't load the internal listfile
const MPQ_OPEN_NO_ATTRIBUTES uint32 = C.MPQ_OPEN_NO_ATTRIBUTES       // Don't open the attributes
const MPQ_OPEN_NO_HEADER_SEARCH uint32 = C.MPQ_OPEN_NO_HEADER_SEARCH // Don't search for the MPQ header past the begin of the file
const MPQ_OPEN_FORCE_MPQ_V1 uint32 = C.MPQ_OPEN_FORCE_MPQ_V1         // Always open the archive as MPQ v 1.00, ignore the "wFormatVersion" variable in the header
const MPQ_OPEN_CHECK_SECTOR_CRC uint32 = C.MPQ_OPEN_CHECK_SECTOR_CRC // On files with MPQ_FILE_SECTOR_CRC, the CRC will be checked when reading file
const MPQ_OPEN_READ_ONLY uint32 = C.MPQ_OPEN_READ_ONLY               // This flag is deprecated. Use STREAM_FLAG_READ_ONLY instead.

// Deprecated
// const MPQ_OPEN_ENCRYPTED uint32 = C.MPQ_OPEN_ENCRYPTED

// Values for SFileOpenFile
const SFILE_OPEN_FROM_MPQ uint32 = C.SFILE_OPEN_FROM_MPQ     // Open the file from the MPQ archive
const SFILE_OPEN_LOCAL_FILE uint32 = C.SFILE_OPEN_LOCAL_FILE // Open a local file

// Flags for SFileAddFile
const MPQ_FILE_IMPLODE uint32 = C.MPQ_FILE_IMPLODE                 // Implode method (By PKWARE Data Compression Library)
const MPQ_FILE_COMPRESS uint32 = C.MPQ_FILE_COMPRESS               // Compress methods (By multiple methods)
const MPQ_FILE_ENCRYPTED uint32 = C.MPQ_FILE_ENCRYPTED             // Indicates whether file is encrypted
const MPQ_FILE_FIX_KEY uint32 = C.MPQ_FILE_FIX_KEY                 // File decryption key has to be fixed
const MPQ_FILE_PATCH_FILE uint32 = C.MPQ_FILE_PATCH_FILE           // The file is a patch file. Raw file data begin with TPatchInfo structure
const MPQ_FILE_SINGLE_UNIT uint32 = C.MPQ_FILE_SINGLE_UNIT         // File is stored as a single unit, rather than split into sectors (Thx, Quantam)
const MPQ_FILE_DELETE_MARKER uint32 = C.MPQ_FILE_DELETE_MARKER     // File is a deletion marker. Used in MPQ patches, indicating that the file no longer exists.
const MPQ_FILE_SECTOR_CRC uint32 = C.MPQ_FILE_SECTOR_CRC           // File has checksums for each sector. Ignored if file is not compressed or imploded.
const MPQ_FILE_SIGNATURE uint32 = C.MPQ_FILE_SIGNATURE             // Present on STANDARD.SNP\(signature). The only occurence ever observed
const MPQ_FILE_EXISTS uint32 = C.MPQ_FILE_EXISTS                   // Set if file exists, reset when the file was deleted
const MPQ_FILE_REPLACEEXISTING uint32 = C.MPQ_FILE_REPLACEEXISTING // Replace when the file exist (SFileAddFile)

// Error codes
const ERROR_SUCCESS uint32 = C.ERROR_SUCCESS
const ERROR_FILE_NOT_FOUND uint32 = C.ERROR_FILE_NOT_FOUND
const ERROR_ACCESS_DENIED uint32 = C.ERROR_ACCESS_DENIED
const ERROR_INVALID_HANDLE uint32 = C.ERROR_INVALID_HANDLE
const ERROR_NOT_ENOUGH_MEMORY uint32 = C.ERROR_NOT_ENOUGH_MEMORY
const ERROR_NOT_SUPPORTED uint32 = C.ERROR_NOT_SUPPORTED
const ERROR_INVALID_PARAMETER uint32 = C.ERROR_INVALID_PARAMETER
const ERROR_NEGATIVE_SEEK uint32 = C.ERROR_NEGATIVE_SEEK
const ERROR_DISK_FULL uint32 = C.ERROR_DISK_FULL
const ERROR_ALREADY_EXISTS uint32 = C.ERROR_ALREADY_EXISTS
const ERROR_INSUFFICIENT_BUFFER uint32 = C.ERROR_INSUFFICIENT_BUFFER
const ERROR_BAD_FORMAT uint32 = C.ERROR_BAD_FORMAT
const ERROR_NO_MORE_FILES uint32 = C.ERROR_NO_MORE_FILES
const ERROR_HANDLE_EOF uint32 = C.ERROR_HANDLE_EOF
const ERROR_CAN_NOT_COMPLETE uint32 = C.ERROR_CAN_NOT_COMPLETE
const ERROR_FILE_CORRUPT uint32 = C.ERROR_FILE_CORRUPT

// Return value for SFileGetFileSize and SFileSetFilePointer
const SFILE_INVALID_SIZE uint32 = C.SFILE_INVALID_SIZE
const SFILE_INVALID_POS uint32 = C.SFILE_INVALID_POS
const SFILE_INVALID_ATTRIBUTES uint32 = C.SFILE_INVALID_ATTRIBUTES

// Move methods for SFileSetFilePointer
const FILE_BEGIN uint32 = C.FILE_BEGIN     // The starting point is 0 (zero) or the beginning of the file.
const FILE_CURRENT uint32 = C.FILE_CURRENT // The starting point is the current file pointer.
const FILE_END uint32 = C.FILE_END         // The starting point is the current end of file.

// Flags for SFileCreateArchive
const MPQ_CREATE_LISTFILE uint32 = C.MPQ_CREATE_LISTFILE     // Also add the (listfile) file
const MPQ_CREATE_ATTRIBUTES uint32 = C.MPQ_CREATE_ATTRIBUTES // Also add the (attributes) file
const MPQ_CREATE_SIGNATURE uint32 = C.MPQ_CREATE_SIGNATURE   // Also add the (signature) file
const MPQ_CREATE_ARCHIVE_V1 uint32 = C.MPQ_CREATE_ARCHIVE_V1 // Creates archive of version 1 (size up to 4GB)
const MPQ_CREATE_ARCHIVE_V2 uint32 = C.MPQ_CREATE_ARCHIVE_V2 // Creates archive of version 2 (larger than 4 GB)
const MPQ_CREATE_ARCHIVE_V3 uint32 = C.MPQ_CREATE_ARCHIVE_V3 // Creates archive of version 3
const MPQ_CREATE_ARCHIVE_V4 uint32 = C.MPQ_CREATE_ARCHIVE_V4 // Creates archive of version 4

// Signature types
const SIGNATURE_TYPE_NONE uint32 = C.SIGNATURE_TYPE_NONE     // The archive has no signature in it
const SIGNATURE_TYPE_WEAK uint32 = C.SIGNATURE_TYPE_WEAK     // The archive has weak signature
const SIGNATURE_TYPE_STRONG uint32 = C.SIGNATURE_TYPE_STRONG // The archive has strong signature

// Flags for SFileVerifyFile
const SFILE_VERIFY_SECTOR_CRC uint32 = C.SFILE_VERIFY_SECTOR_CRC // Verify sector checksum for the file, if available
const SFILE_VERIFY_FILE_CRC uint32 = C.SFILE_VERIFY_FILE_CRC     // Verify file CRC, if available
const SFILE_VERIFY_FILE_MD5 uint32 = C.SFILE_VERIFY_FILE_MD5     // Verify file MD5, if available
const SFILE_VERIFY_RAW_MD5 uint32 = C.SFILE_VERIFY_RAW_MD5       // Verify raw file MD5, if available
const SFILE_VERIFY_ALL uint32 = C.SFILE_VERIFY_ALL               // Verify every checksum possible

// Return values for SFileVerifyFile
const VERIFY_OPEN_ERROR uint32 = C.VERIFY_OPEN_ERROR                       // Failed to open the file
const VERIFY_READ_ERROR uint32 = C.VERIFY_READ_ERROR                       // Failed to read all data from the file
const VERIFY_FILE_HAS_SECTOR_CRC uint32 = C.VERIFY_FILE_HAS_SECTOR_CRC     // File has sector CRC
const VERIFY_FILE_SECTOR_CRC_ERROR uint32 = C.VERIFY_FILE_SECTOR_CRC_ERROR // Sector CRC check failed
const VERIFY_FILE_HAS_CHECKSUM uint32 = C.VERIFY_FILE_HAS_CHECKSUM         // File has CRC32
const VERIFY_FILE_CHECKSUM_ERROR uint32 = C.VERIFY_FILE_CHECKSUM_ERROR     // CRC32 check failed
const VERIFY_FILE_HAS_MD5 uint32 = C.VERIFY_FILE_HAS_MD5                   // File has data MD5
const VERIFY_FILE_MD5_ERROR uint32 = C.VERIFY_FILE_MD5_ERROR               // MD5 check failed
const VERIFY_FILE_HAS_RAW_MD5 uint32 = C.VERIFY_FILE_HAS_RAW_MD5           // File has raw data MD5
const VERIFY_FILE_RAW_MD5_ERROR uint32 = C.VERIFY_FILE_RAW_MD5_ERROR       // Raw MD5 check failed

// Return values for SFileVerifyArchive
const ERROR_NO_SIGNATURE uint32 = C.ERROR_NO_SIGNATURE                     // There is no signature in the MPQ
const ERROR_VERIFY_FAILED uint32 = C.ERROR_VERIFY_FAILED                   // There was an error during verifying signature (like no memory)
const ERROR_WEAK_SIGNATURE_OK uint32 = C.ERROR_WEAK_SIGNATURE_OK           // There is a weak signature and sign check passed
const ERROR_WEAK_SIGNATURE_ERROR uint32 = C.ERROR_WEAK_SIGNATURE_ERROR     // There is a weak signature but sign check failed
const ERROR_STRONG_SIGNATURE_OK uint32 = C.ERROR_STRONG_SIGNATURE_OK       // There is a strong signature and sign check passed
const ERROR_STRONG_SIGNATURE_ERROR uint32 = C.ERROR_STRONG_SIGNATURE_ERROR // There is a strong signature but sign check failed
