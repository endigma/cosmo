import { docsBaseURL } from "@/lib/constants";
import { AuditLog } from "@wundergraph/cosmo-connect/dist/platform/v1/platform_pb";
import { EmptyState } from "./empty-state";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { formatDateTime } from "@/lib/format-date";
import { capitalize } from "@/lib/utils";
import { pascalCase } from "change-case";
import { AiOutlineAudit } from "react-icons/ai";
import { PiKeyBold, PiRobotFill, PiUserBold } from "react-icons/pi";

export const Empty = (params: { unauthorized: boolean }) => {
  if (params.unauthorized) {
    return (
      <EmptyState
        title="Unauthorized"
        description="You are not authorized to manage this organization."
      />
    );
  }

  return (
    <EmptyState
      icon={<AiOutlineAudit />}
      title={params.unauthorized ? "Unauthorized" : "No audit logs"}
      description={
        <div className="space-x-1">
          <span>You are not authorized to view audit logs.</span>
          <a
            target="_blank"
            rel="noreferrer"
            href={docsBaseURL + "/studio/audit-log"}
            className="text-primary"
          >
            Learn more.
          </a>
        </div>
      }
    />
  );
};

export const AuditLogTable = ({ logs }: { logs?: AuditLog[] }) => {
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead className="px-4">Actor</TableHead>
          <TableHead className="px-4">Action</TableHead>
          <TableHead className="px-4">Date</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {logs?.map(
          ({
            id,
            actorDisplayName,
            actorType,
            auditAction,
            createdAt,
            action,
            auditableDisplayName,
            targetDisplayName,
            targetType,
          }) => {
            let preParagraph = null;
            let postParagraph = null;

            if (auditAction === "organization_invitation.created") {
              postParagraph = "for";
            } else if (auditAction === "member_role.updated") {
              preParagraph = "role for";
              postParagraph = "to";
            } else if (auditableDisplayName) {
              preParagraph = "in";
            }

            let label = null;

            if (targetDisplayName) {
              label = (
                <>
                  {preParagraph && (
                    <span className="text-gray-500 dark:text-gray-400">
                      {preParagraph}
                    </span>
                  )}

                  <span
                    className="inline-block max-w-md truncate"
                    title={pascalCase(targetType)}
                  >
                    <span className="text-purple whitespace-nowrap font-mono">
                      {targetDisplayName}
                    </span>
                  </span>
                </>
              );
            }

            const actionView = (
              <>
                <span className="text-gray-500 dark:text-gray-400">
                  {capitalize(action)}
                </span>
                {label}
                {auditableDisplayName && (
                  <>
                    {postParagraph && (
                      <span className="text-gray-500 dark:text-gray-400">
                        {postParagraph}
                      </span>
                    )}
                    <span className="inline-block max-w-md truncate text-primary">
                      {auditableDisplayName}
                    </span>
                  </>
                )}
              </>
            );
            return (
              <TableRow
                key={id}
                className="group py-1 even:bg-secondary/20 hover:bg-secondary/40"
              >
                <TableCell className="flex px-4 font-medium">
                  <span className="flex items-center justify-center space-x-2">
                    {actorType === "api_key" && (
                      <PiKeyBold className="h-4 w-4" title="API Key activity" />
                    )}
                    {actorType === "user" && (
                      <PiUserBold className="h-4 w-4" title="User activity" />
                    )}
                    {actorType === "system" && (
                      <PiRobotFill
                        className="h-4 w-4"
                        title="System activity"
                      />
                    )}
                    <span className="block font-medium">
                      {actorDisplayName}
                    </span>
                  </span>
                </TableCell>
                <TableCell className="px-4 font-medium">
                  <div className="justify-center space-y-2">
                    <div className="flex flex-wrap space-x-1.5">
                      {actionView}
                    </div>
                    <div className="focus-ring flex inline-flex items-center rounded-full border px-1.5 font-mono text-sm text-gray-500 dark:text-gray-200">
                      {auditAction}
                    </div>
                  </div>
                </TableCell>
                <TableCell className="px-4 font-medium">
                  {formatDateTime(new Date(createdAt))}
                </TableCell>
              </TableRow>
            );
          },
        )}
      </TableBody>
    </Table>
  );
};
